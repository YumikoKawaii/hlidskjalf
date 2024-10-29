from locust import HttpUser, task, between
import random
import csv

class User(HttpUser):
    wait_time = between(1, 5)
    accounts = []

    @classmethod
    def on_start_class(cls):
        with open("../etl/faker/data/accounts.csv", newline='' ) as file:
            reader = csv.DictReader(file)
            cls.accounts = [{"email": row["email"], "password": row["password"]} for row in reader]

    def on_start(self):
        account = random.choice(self.accounts)
        email = account["email"]
        password = account["password"]
        response = self.client.post("http://127.0.0.1:10080/api/v1/login", json={"email": email, "password": password})
        self.token = response.json().get("token")

    @task
    def test_get_post_detail(self):
        headers = {
            "Authorization": self.token
        }
        self.client.get("/api/v1/posts-detail", headers=headers)
