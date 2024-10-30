from locust import HttpUser, task, between

class MyUser(HttpUser):
    wait_time = between(1, 5)
    total_requests = 1000
    request_counter = 0

    @task
    def hello_world(self):
        if self.request_counter > self.total_requests:
            self.stop()

        self.client.get("/api/v1/quaternary")
        self.request_counter += 1
