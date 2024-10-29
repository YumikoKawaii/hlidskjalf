from locust import HttpUser, task, between
import requests

class MyUser(HttpUser):
    wait_time = between(1, 5)

    @task
    def test_quaternary(self):
        self.client.get("/api/v1/quaternary")

    # @task
    # def test_faux(self):
    #     self.client.get("/api/v1/faux")
    #
    # @task
    # def test_fictio(self):
    #     self.client.get("/api/v1/fictio")

