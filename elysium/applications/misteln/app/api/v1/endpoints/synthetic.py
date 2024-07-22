from fastapi import APIRouter
from applications.misteln.app.dto.pawn.synthetic import GetSyntheticStudentsResponse
from applications.misteln.app.adapter.pawn import PawnAdapter

router = APIRouter()


@router.get('', response_model=GetSyntheticStudentsResponse)
def get_synthetic_students():
    data = PawnAdapter.get_synthetic_students()
    return data