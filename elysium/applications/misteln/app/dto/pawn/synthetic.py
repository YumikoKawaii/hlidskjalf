from typing import List, Optional

from applications.misteln.app.dto.base import CamelBaseModel


class SyntheticStudent(CamelBaseModel):
    id: str
    full_name: Optional[str] = None
    age: int
    sex: str
    major: str
    level: str
    gpa: float
    hobbies: List[str]
    country: str
    province: str


class GetSyntheticStudentsResponse(CamelBaseModel):
    code: int
    message: str
    students: List[SyntheticStudent]