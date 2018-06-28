# Data Service

REST API's that return current UTC time to simulate services that serve data to the user interface.

*Load Balanced Data Service*: for testing traffic routing and canary deployments

*Rate Limited Data Service*: for testing rate limiting of backend services

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_data_svc
    $ export IMAGE_TAG=0.1
    $ make
```

## Deployment Instructions

Edit manifests to use your image if you built your own.  Or you can use the one already there.
```
    $ kubectl apply -f lb_data_svc.yaml
    $ kubectl apply -f rate_limit_data_service.yaml
```

