from fastapi import FastAPI
from setting import settings
from app.api.v1.endpoints import v1_router


def initialize() -> FastAPI:
    application = FastAPI(
        title=settings.PROJECT_TITLE,
        docs_url=''
    )
    application.include_router(prefix='/api', router=v1_router)

    return application