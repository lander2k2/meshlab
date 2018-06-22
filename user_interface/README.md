# User Interface

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_ui
    $ export IMAGE_TAG=0.1
    $ make
```

## Deployment Instructions

Edit ui.yaml to use your own image if you built your own.  Or you can use the one already there.

```
    $ kubectl apply -f ui.yaml
```

