# User Interface

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_ui
    $ export IMAGE_TAG=0.3
    $ make
```

## Deployment Instructions

Edit ui.yaml to use your own image if you built your own.  Or you can use the one already there.

```
    $ kubectl apply -f ui.yaml
```

If you uncomment the line `command: ["/ui", "no-canary"]` the user interface will not make calls to the canary back-end which is useful for testing failure scenarios.

