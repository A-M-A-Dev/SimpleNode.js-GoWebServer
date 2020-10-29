import time
from locust import HttpUser, task, between

class NodeWriteUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def getNodeWrite(self):
        self.client.get("/helloworld/node/write?line=68")