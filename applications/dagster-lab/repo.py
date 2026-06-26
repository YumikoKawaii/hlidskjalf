import random
import time

import dagster as dg
from dagster_celery_k8s import celery_k8s_job_executor


@dg.asset
def fetched(context: dg.AssetExecutionContext) -> list[int]:
    """Pretend to fetch some records."""
    n = random.randint(3, 8)
    context.log.info(f"fetched {n} records")
    time.sleep(2)
    return list(range(n))


@dg.asset
def processed(context: dg.AssetExecutionContext, fetched: list[int]) -> list[int]:
    """Square each record."""
    out = [r * r for r in fetched]
    context.log.info(f"processed {len(out)} records -> {out}")
    time.sleep(2)
    return out


@dg.asset
def stored(context: dg.AssetExecutionContext, processed: list[int]) -> None:
    """Pretend to persist results."""
    context.log.info(f"stored {len(processed)} records, sum={sum(processed)}")
    time.sleep(2)


# Materialize the whole asset graph as one run, each step a Celery -> K8s Job.
etl_job = dg.define_asset_job(
    name="etl_job",
    selection=dg.AssetSelection.all(),
    executor_def=celery_k8s_job_executor,
)

defs = dg.Definitions(
    assets=[fetched, processed, stored],
    jobs=[etl_job],
)
