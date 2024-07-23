from fastapi import APIRouter
from . import synthetic

v1_router = APIRouter()

v1_router.include_router(synthetic.router, prefix="/synthetic")