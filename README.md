# Meshlab

A collection of components to test features and develop solutions in an Istio service mesh.  Tested with Istio v0.8.

Assuming you have a running Kubernetes cluster with Istio installed and automatic sidecar injection:

1. Get the ELB DNS name for the Istion ingress gateway:

```
    $ kubectl get svc -n istio-system istio-ingressgateway -o wide | awk '{ print $4 }' | grep amazonaws
```

2. Add that DNS name to the `meshlab-ext` service entry in `meshlab_gateway.yaml`.  This will allow the traffic generator to egress from the cluster and back into the UI.

3. Enable automatic sidecar injection:

```
    $ kubectl label ns default istio-injection=enabled
```

4. Deploy the meshlab gateway, virtual service and service entry.

```
    $ kubectl apply -f meshlab_gateway.yaml
```

5. See the instructions in `./user_interface`, `./data_service` and `./traffic_generator` for further steps on deploying those components..

