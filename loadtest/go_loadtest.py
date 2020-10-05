import time
from locust import HttpUser, task, between

class GoUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def getGo(self):
        self.client.get(f"/helloworld/go")