# Data Service

REST API's to simulate data services that return data to the user interface.

* `load_balanced`: for basic testing of a back-end service

* `rate_limited`: for testing rate limiting

* `canary`: for prototyping canary deployment roll-outs

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_data_svc
    $ export IMAGE_TAG=$(cat VERSION)
    $ make
```

## Deployment Instructions

Simply `kubectl apply` the manifests for `load_balanced` or `rate_limited` data services.  See below for `canary`.

## Canary Deployment

### Concept

These scripts and manifests are intended as a prototype for an operator that can manage canary deployments using istio primitives.  This proof-of-concept involves error checking for the following:

* excessive 500 errors from the upgraded version of the service
* poor response performance from the upgraded version
* a significant drop in client requests to the back-end service

If any of these tests fail, the roll-out is halted and reverted back to the original version.

The various `canary_data_service_*` manifests allow various degrees of traffic shifting from v0.1 to v0.2 in a simulated canary deployment.

Manifest | v0.1 | v0.2.x
---|:---:|:---:
`canary_data_service_0.yaml` | 100% | 0%
`canary_data_service_3.yaml` | 97% | 3%
`canary_data_service_20.yaml` | 80% | 20%
`canary_data_service_50.yaml` | 50% | 50%
`canary_data_service_100.yaml` | 0% | 100%

### Simulate error threshold violation

How to:

1. Tell the traffic generator to spawn 200 users and spawn at 50 per second to bring them all up in a few seconds.  Each user will make a request to the user interface approx every 1 second.  The UI will, in turn, make requests to the back-end services.  The traffic generator will ensure enough requests are made to violate the error threshold.

2. Roll out the canary deployment:
```
    $ cd canary
    $ ./canary-coalmine.sh breaking_error v0.2.0
```

3. After that starts, in another shell session, port forward the prometheus pod to query metrics:
```
    $ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090
```

What happens:

1. Version 0.1 is deployed to simulate the existing deployment of a service.  A count-down begins to allow the initial version to spin up.

2. At the end of the count-down, a single pod for version 0.2.0 is deployed and 3% of traffic is routed to this version.  The other 97% of requests still route to v0.1.  A count-down will begin to allow v0.2.0 to spin up and receive enough requests to exceed the error threshold.

3. An error check will be run.  This is performed by collecting metrics from Prometheus.  It will learn that more than 10 500's were received by the UI from the data service and will abort the deployment, reverting all traffic back to v0.1 and deleting the v0.2.0 deployment.

### Simulate performance threshold violation

How to:

1. Tell the traffic generator to spawn 200 users and spawn at 50 per second to bring them all up in a few seconds.  Each user will make a request to the user interface approx every 1 second.  The UI will, in turn, make requests to the back-end services.  The traffic generator will ensure enough requests are made to violate the error threshold.

2. Roll out the canary deployment:
```
    $ cd canary
    $ ./canary-coalmine.sh breaking_performance v0.2.1
```

3. After that starts, in another shell session, port forward the prometheus pod to query metrics:
```
    $ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090
```

What happens:

1. Version 0.1 is deployed to simulate the existing deployment of a service.  A count-down begins to allow the initial version to spin up.

2. At the end of the count-down, a single pod for version 0.2.1 is deployed and 3% of traffic is routed to this version.  The other 97% of requests still route to v0.1.  A count-down will begin to allow v0.2.1 to spin up and receive enough requests to test the performance threshold.

3. A performance check will be run.  This is performed by collecting metrics from Prometheus.  It will learn that all responses from the v0.2.1 back-end service to the UI took longer than 1 second.

4. With this result, all traffic is reverted back to v0.1 and v0.2.1 deployment is deleted.

### Simulate client request count violation

How to:

1. Instruct the user interface to stop sending requests to the canary deployment back-end.  Do this by uncommenting the line `command: ["/ui", "no-canary"]` in the UI's manifest.  Then `kubectl apply -f` the manifest file.

2. Tell the traffic generator to spawn 200 users and spawn at 50 per second to bring them all up in a few seconds.  Each user will make a request to the user interface approx every 1 second.  The UI will, in turn, make requests to the back-end services.  The traffic generator will ensure enough requests are made to violate the error threshold.

3. Roll out the canary deployment:
```
    $ cd canary
    $ ./canary-coalmine.sh working v0.2.2
```

4. After that starts, in another shell session, port forward the prometheus pod to query metrics:
```
    $ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090
```

What happens:

1. Version 0.1 is deployed to simulate the existing deployment of a service.  A count-down begins to allow the initial version to spin up.  If you already have the v0.1 still running from a previous simulation, nothing will change.

2. At the end of the count-down, a single pod for version 0.2.2 is deployed and 3% of traffic is routed to this version.  The other 97% of requests still route to v0.1.  A count-down will begin to allow v0.2.2 to spin up and receive enough requests to test the request count from UI to back-end service.

3. A requets check will be run.  This is performed by collecting metrics from Prometheus.  It will learn that less than the aribitrary threshold of 10 requests have been made by UI to the back-end service.  In fact, it will be zero requests because we told the UI specifically to stop send requests to that service altogether.

4. Due to the requests check failure, all traffic is reverted back v0.1 and the v0.2.2 deployment is deleted.

### Simulate successful canary deployment

How to:

1. Roll out the canary deployment:
```
    $ cd canary
    $ ./canary-coalmine.sh working v0.2.2
```

2. After that starts, in another shell session, port forward the prometheus pod to query metrics:
```
    $ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090
```

What happens:

1. Version 0.1 is deployed to simulate the existing deployment of a service.  A count-down begins to allow the initial version to spin up.  If you already have the v0.1 still running from a previous simulation, nothing will change.

2. At the end of the count-down, a single pod for version 0.2.2 is deployed and 3% of traffic is routed to this version.  The other 97% of requests still route to v0.1.  A count-down will begin to allow v0.2.2 to spin up and receive enough requests to exceed the error threshold.

3. An error check will be run.  This is performed by collecting metrics from Prometheus.  In this case, the error check will pass on v0.2.2 and the roll-out will continue.

4. The traffic routed to v0.2.2 will be increased to 20%.  A second pod for the new version will spin up to accommodate the additional traffic.  Then another error check will run and pass.

5. The traffic routed to v0.2.2 will be increased to 50%.  The replicas to v0.1 will begin to scale back.  Replicas for v0.2.2 will be scaled out further.  Another error check will run and pass.

6. All traffic will be routed to the new version.

7. Finally the v0.1 deployment will be deleted.

