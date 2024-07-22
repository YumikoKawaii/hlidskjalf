from app.dto.pawn.synthetic import GetSyntheticStudentsResponse
from app.adapter.pawn import PawnAdapter


class PawnService:

    @classmethod
    def get_synthetic_students(cls) -> GetSyntheticStudentsResponse:
        resp_data = PawnAdapter.get_synthetic_students()
        print('something here')
        return GetSyntheticStudentsResponse.parse_obj(resp_data)
