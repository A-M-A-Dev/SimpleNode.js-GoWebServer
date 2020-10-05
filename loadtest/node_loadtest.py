import time
from locust import HttpUser, task, between

class NodeUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def getNode(self):
        self.client.get(f"/helloworld/node")