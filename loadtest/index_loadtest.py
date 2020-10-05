import time
from locust import HttpUser, task, between

class HelloworldUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def getHelloworld(self):
        self.client.get(f"/helloworld")