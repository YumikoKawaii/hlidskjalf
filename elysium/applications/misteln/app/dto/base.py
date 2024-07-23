from typing import Optional, List, Union, Dict

from pydantic import BaseModel

from app.helper.utilities import convert_str_to_camel


class CamelBaseModel(BaseModel):
    class Config:
        alias_generator = convert_str_to_camel
        allow_population_by_field_name = True
