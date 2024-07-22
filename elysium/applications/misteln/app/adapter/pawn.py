from typing import Dict
import requests
import logging

from setting import settings
from app.dto.pawn.synthetic import GetSyntheticStudentsResponse


class PawnAdapter:
    _host = settings.PAWN_SERVICE_ADDRESS

    @classmethod
    def get_headers(cls) -> Dict:
        return {'Content-Type': 'application/json', 'access-token': ''}

    @classmethod
    def call(cls, method: str, path: str, params=None, payload=None, timeout=30):
        url = cls._host + path
        headers = cls.get_headers()
        try:
            resp: requests.Response = requests.request(method, url, params=params, json=payload,headers=headers,timeout=timeout)
            r = resp.json()
            if resp.status_code != 200 or (resp.status_code == 200 and resp.json().get('code') != 200):
                logging.warning('Calling Pawn got unexpected result, url: %s, params: %s, payload: %s, response: %s',
                                url, str(params), str(payload), str(resp.status_code), resp.text)
            return resp
        except Exception as e:
            logging.error(f'Error {e}')
            raise e

    @classmethod
    def get_synthetic_students(cls) -> GetSyntheticStudentsResponse:
        try:
            resp = cls.call(method='GET',
                            path=f'/api/v1/students',
                            params=None,
                            payload=None,
                            )
            if resp.status_code == 200 and resp.json()['code'] == 200:
                return GetSyntheticStudentsResponse.parse_obj(resp.json())
        except Exception as e:
            raise e
