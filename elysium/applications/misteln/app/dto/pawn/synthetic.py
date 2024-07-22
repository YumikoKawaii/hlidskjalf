from typing import List

from applications.misteln.app.dto.base import CamelBaseModel


class SyntheticStudent(CamelBaseModel):
    id: str
    name: str
    age: int
    sex: str
    major: str
    level: str
    gpa: float
    hobbies: List[str]
    country: str
    province: str


class GetSyntheticStudentsResponse(CamelBaseModel):
    code: str
    message: str
    students: List[SyntheticStudent]