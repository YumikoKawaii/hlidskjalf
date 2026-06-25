import random
import time

import dagster as dg
from dagster_celery_k8s import celery_k8s_job_executor


@dg.op
def fetch(context: dg.OpExecutionContext) -> list[int]:
    """Pretend to fetch some records."""
    n = random.randint(3, 8)
    context.log.info(f"fetched {n} records")
    time.sleep(2)
    return list(range(n))


@dg.op
def process(context: dg.OpExecutionContext, records: list[int]) -> list[int]:
    """Square each record."""
    out = [r * r for r in records]
    context.log.info(f"processed {len(out)} records -> {out}")
    time.sleep(2)
    return out


@dg.op
def store(context: dg.OpExecutionContext, records: list[int]) -> None:
    """Pretend to persist results."""
    context.log.info(f"stored {len(records)} records, sum={sum(records)}")
    time.sleep(2)


@dg.job(executor_def=celery_k8s_job_executor)
def etl_job():
    store(process(fetch()))


defs = dg.Definitions(jobs=[etl_job])
