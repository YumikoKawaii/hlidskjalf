from typing import Union

from fastapi import FastAPI

application = FastAPI()


@application.get("/api/v1/greet")
def greet():
    return {"Hello"}
