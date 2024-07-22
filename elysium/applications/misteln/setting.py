import os

from dotenv import load_dotenv
from pydantic_settings import BaseSettings


load_dotenv(verbose=True)


class Settings(BaseSettings):
    PROJECT_TITLE: str = 'MISTELN'
    PAWN_SERVICE_ADDRESS: str = os.getenv('PAWN_SERVICE_ADDRESS','http://127.0.0.1:8080')


settings = Settings()