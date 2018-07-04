# Traffic Generator

## Build Instructions

```
    $ export IMAGE_REPO=quay.io/myimages/meshlab_ui
    $ export IMAGE_TAG=0.1
    $ make
```

## Deployment Instructions

1. Get the ELB DNS name for the Istion ingress gateway:

```
    $ kubectl get svc -n istio-system istio-ingressgateway -o wide | awk '{ print $4 }' | grep amazonaws
```

2. Add that DNS name to the `meshlab-traffic-config` configmap in `traffic_generator.yaml`.  This will tell locust where to send requests when generating traffic.

3. Replace the image name in the manifest if using your own build.

4. Deploy

```
    $ kubectl apply -f traffic_generator.yaml
```

