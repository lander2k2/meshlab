# Data Service

A simple REST API that just serves current UTC time for initial testing.

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_data_svc
    $ export IMAGE_TAG=0.1
    $ make
```

## Deployment Instructions

Edit data_svc.yaml to use your own image if you built your own.  Or you can use the one already there.

```
    $ kubectl apply -f data_svc.yaml
```

