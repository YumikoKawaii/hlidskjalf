import os

from dotenv import load_dotenv
from pydantic import BaseSettings


load_dotenv(verbose=True)


class Settings(BaseSettings):
    PAWN_SERVICE_ADDRESS: str = os.getenv('PAWN_SERVICE_ADDRESS','localhost:9090')
