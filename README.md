# Meshlab

A collection of components to test features and develop solutions in an Istio service mesh.

Assuming you have a running Kubernetes cluster with Istio installed and automatic sidecar injection, start here:
```
    $ kubectl label ns default istio-injection=enabled
    $ kubectl apply -f meshlab-gateway.yaml
```

Then see `user_interface` and `data_service` for further steps.

