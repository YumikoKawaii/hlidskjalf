from locust import HttpUser, task, between
import requests

class MyUser(HttpUser):
    wait_time = between(1, 5)

    @task
    def test_endpoint_with_new_connection(self):
        with requests.Session() as session:
            session.get(f"{self.environment.host}/api/v1/echo")
