#!/bin/bash

observe() {
    OBSERVE_SECONDS=$1
    while [ $OBSERVE_SECONDS -ge 0 ]; do
        echo $OBSERVE_SECONDS...
        OBSERVE_SECONDS=$((OBSERVE_SECONDS-1))
        sleep 1
    done
}

echo "Initialize!"

echo "Deploy v1"
kubectl apply -f canary_data_service_0.yaml

echo "Observe v1 deployment..."
observe 60

echo "Send the canary down the coalmine!"
kubectl apply -f canary_data_service_3.yaml

echo "Observe the canary..."
observe 30

echo "Canary is still chirping - send a few miners!"
kubectl apply -f canary_data_service_20.yaml

echo "Observe the miners..."
observe 30

echo "Miners are still breathing - send some more!"
kubectl apply -f canary_data_service_50.yaml

echo "Observe the additional miners..."
observe 30

echo "All looks good - send 'em all!"
kubectl apply -f canary_data_service_100.yaml

echo "Observe the glory..."
observe 30

echo "Clean up"
kubectl delete deploy meshlab-canary-data-svc-v1

