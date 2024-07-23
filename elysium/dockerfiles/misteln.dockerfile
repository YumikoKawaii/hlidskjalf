FROM python:3.8.19-slim AS builder

WORKDIR /app

COPY ./elysium/requirements.txt .

RUN pip install -r requirements.txt

COPY ./elysium/applications/misteln .

EXPOSE 8000

#CMD ["bash", "uvicorn main:application", "--reload"]
CMD ["uvicorn", "main:application"]
#CMD ["bash", "-c","while true; do echo '1'; sleep 1; done"]