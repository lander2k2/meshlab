from locust import HttpLocust, TaskSet

def index(l):
    l.client.get("/ui")

class Behavior(TaskSet):
    tasks = {index:1}

class User(HttpLocust):
    task_set = Behavior
    min_wait = 500
    max_wait = 1500

