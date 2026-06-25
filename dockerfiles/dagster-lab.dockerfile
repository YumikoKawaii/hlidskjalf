FROM python:3.12-slim

ENV PYTHONUNBUFFERED=1 \
    DAGSTER_HOME=/opt/dagster/dagster_home

WORKDIR /opt/dagster/app

COPY applications/dagster-lab/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY applications/dagster-lab/ ./

RUN mkdir -p $DAGSTER_HOME

EXPOSE 3030

# Code location server; the Dagster Helm chart overrides args per deployment.
CMD ["dagster", "api", "grpc", "-h", "0.0.0.0", "-p", "3030", "-f", "repo.py"]
