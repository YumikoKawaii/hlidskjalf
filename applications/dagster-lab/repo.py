"""Partition-based fan-out lab, mirroring the integration-ingress pattern.

Fan-out = one RUN per dynamic partition (not op-level .map()). A sensor polls a
config source, registers a partition per item, and emits one RunRequest each.
A dynamically-partitioned asset does the work for a single partition, rate-limited
via a concurrency_key tag. Execution is distributed over Celery -> K8s Jobs.
"""

import random
import time

import os

import dagster as dg
from dagster_celery import celery_executor

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
    """Process exactly ONE shop (one partition = one run)."""
    shop_id = context.partition_key
    context.log.info(f"syncing {shop_id}")
    time.sleep(2)
    total = sum(ord(c) for c in shop_id)
    context.log.info(f"{shop_id} done, total={total}")
    return {"shop_id": shop_id, "total": total}


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
