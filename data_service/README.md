# Data Service

REST API's that return current UTC time to simulate services that serve data to the user interface.

*Load Balanced Data Service*: for testing traffic routing and canary deployments

*Rate Limited Data Service*: for testing rate limiting of backend services

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_data_svc
    $ export IMAGE_TAG=$(cat VERSION)
    $ make
```

## Deployment Instructions

Edit manifests to use your image if you built your own.  Or you can use the one already there.
```
    $ kubectl apply -f lb_data_svc.yaml
    $ kubectl apply -f rate_limit_data_service.yaml
```

## Canary Deployment

The various `canary_data_service_*` manifests allow various degrees of traffic shifting from v1 to v2 in a simulated canary deployment.

Manifest | v1 | v2
---|:---:|:---:
`canary_data_service_0.yaml` | 100% | 0%
`canary_data_service_3.yaml` | 97% | 3%
`canary_data_service_20.yaml` | 80% | 20%
`canary_data_service_50.yaml` | 50% | 50%
`canary_data_service_100.yaml` | 0% | 100%

Apply these manifests manually or watch them roll out with:
```
    $ ./canary-coalmine.sh
```

