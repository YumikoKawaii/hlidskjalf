"""Partition-based fan-out lab, mirroring the integration-ingress pattern.

Fan-out = one RUN per dynamic partition (not op-level .map()). A sensor polls a
config source, registers a partition per item, and emits one RunRequest each.
A dynamically-partitioned asset does the work for a single partition, rate-limited
via a concurrency_key tag. Execution is distributed over Celery -> K8s Jobs.

Also reproduces the SPI flag false-OFF bug: a thread-backed FlagClient built
ONCE in the parent (module import) is reused inside the forked Celery
ForkPoolWorker child; its background thread does not survive os.fork(), so
fetch_v2() returns empty -> fail-safe default reports the flag OFF for an ON
shop. See applications/dagster-lab/flag_client.py and the ticket
spi-flag-false-off-celery-fork-amplitude-client.
"""

import os
import random
import time

import dagster as dg
from dagster_celery import celery_executor

from flag_client import FlagClient

FLAG_KEY = "shoppayinstallment_enable_20260130"

# Built ONCE in the parent at import — same as FeatureFlagResource building
# self._client in setup_for_execution. Celery prefork children inherit this
# object via os.fork(), but its background thread dies at the fork.
_PARENT_CLIENT = FlagClient()


def is_enabled(client: FlagClient, key: str) -> bool:
    """Fail-safe-default rule from the real code: empty answer => OFF."""
    variants = client.fetch_v2()
    if not variants:
        # Empty (thread dead after fork) is indistinguishable here from a
        # genuinely-off flag -> we skip the shop with no error. The bug.
        return False
    return variants.get(key) == "on"

# Celery executor (broker/backend from env), mirroring integration-ingress.
# Run launcher (K8sRunLauncher) gives one pod per run; this executor dispatches
# the run's steps as Celery tasks to the worker pool over Redis.
configured_celery_executor = celery_executor.configured(
    {
        "broker": os.getenv("DAGSTER_CELERY_BROKER_URL", "redis://redis:6379/0"),
        "backend": os.getenv("DAGSTER_CELERY_BACKEND_URL", "redis://redis:6379/1"),
        "config_source": {
            "task_create_missing_queues": True,
            "task_default_queue": "dagster",
        },
    }
)

# One partition per "shop" — partitions are added at runtime by the sensor,
# exactly like integration-ingress's tenant_partitions.
shop_partitions = dg.DynamicPartitionsDefinition(name="shop_id")


def _list_shops() -> list[str]:
    """Stand-in for the MongoDB shop list in integration-ingress.
    Returns a runtime-varying set of shop ids to fan out over."""
    n = random.randint(3, 8)
    return [f"shop_{i}" for i in range(n)]


@dg.asset(
    partitions_def=shop_partitions,
    group_name="fanout_lab",
    op_tags={"concurrency_key": "shop_sync"},  # rate-limited via tag_concurrency_limits
    retry_policy=dg.RetryPolicy(max_retries=2, delay=10, backoff=dg.Backoff.EXPONENTIAL),
)
def shop_sync(context: dg.AssetExecutionContext) -> dict:
    """Process exactly ONE shop (one partition = one run).

    Gated by the SPI flag via the parent-built client. When this step runs in a
    forked Celery ForkPoolWorker child, the inherited client's background thread
    is dead -> fetch_v2() empty -> is_enabled False -> 'flag OFF, skipping',
    even though the flag is ON. PID/thread fields are logged as evidence."""
    shop_id = context.partition_key

    enabled = is_enabled(_PARENT_CLIENT, FLAG_KEY)
    context.log.info(
        f"flag check {shop_id}: enabled={enabled} "
        f"pid={os.getpid()} client_owner_pid={_PARENT_CLIENT._owner_pid} "
        f"thread_alive_here={_PARENT_CLIENT.thread_alive_here()}"
    )
    if not enabled:
        context.log.info(f"Flag {FLAG_KEY} OFF — skipping {shop_id}")
        return {"shop_id": shop_id, "skipped": True}

    context.log.info(f"syncing {shop_id}")
    time.sleep(2)
    total = sum(ord(c) for c in shop_id)
    context.log.info(f"{shop_id} done, total={total}")
    return {"shop_id": shop_id, "total": total, "skipped": False}


# One K8s pod per run (K8sRunLauncher); steps dispatched as Celery tasks to the
# worker pool over Redis — mirrors integration-ingress.
shop_sync_job = dg.define_asset_job(
    name="shop_sync_job",
    selection=dg.AssetSelection.assets(shop_sync),
    executor_def=configured_celery_executor,
)


@dg.sensor(
    name="shop_sync_sensor",
    job=shop_sync_job,
    minimum_interval_seconds=30,
    default_status=dg.DefaultSensorStatus.STOPPED,
)
def shop_sync_sensor(context: dg.SensorEvaluationContext) -> dg.SensorResult:
    """Poll the shop list, register a partition per new shop, emit one
    RunRequest per shop -> FAN-OUT into N parallel partitioned runs."""
    shops = _list_shops()

    existing = set(context.instance.get_dynamic_partitions(shop_partitions.name))
    new_partitions = [s for s in shops if s not in existing]

    run_requests = [
        dg.RunRequest(
            partition_key=shop_id,
            run_key=f"{shop_id}_{context.cursor or '0'}",
            tags={"dagster/partition": shop_id, "shop_id": shop_id},
        )
        for shop_id in shops
    ]

    context.log.info(
        f"sensor: {len(run_requests)} runs, {len(new_partitions)} new partitions"
    )

    return dg.SensorResult(
        run_requests=run_requests,
        dynamic_partitions_requests=[
            shop_partitions.build_add_request(new_partitions)
        ]
        if new_partitions
        else [],
    )


defs = dg.Definitions(
    assets=[shop_sync],
    jobs=[shop_sync_job],
    sensors=[shop_sync_sensor],
)
