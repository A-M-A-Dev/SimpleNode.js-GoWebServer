import time
from locust import HttpUser, task, between


class NodeAdderUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def postNodeAdder(self):
        self.client.post('/helloworld/node/adder', json={"a": 2, "b": 3})
