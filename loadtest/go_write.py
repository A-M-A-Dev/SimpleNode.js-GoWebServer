import time
from locust import HttpUser, task, between

class GoWriteUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def getGoWrite(self):
        self.client.get("/helloworld/go/write?line=68")