import time
from locust import HttpUser, task, between


class GoAdderUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def postGoAdder(self):
        self.client.post('/helloworld/go/adder', json={"a": 2, "b": 3})
