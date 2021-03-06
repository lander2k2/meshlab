#!/bin/bash

MODE=$1
VERSION=$2

DEPLOY_NAME=$(echo "meshlab-canary-data-svc-$VERSION" | tr "." "-")

observe() {
    OBSERVE_SECONDS=$1
    while [ $OBSERVE_SECONDS -ge 0 ]; do
        echo $OBSERVE_SECONDS...
        OBSERVE_SECONDS=$((OBSERVE_SECONDS-1))
        sleep 1
    done
}

deploy_check() {
    ./error-check.sh $VERSION
    if [ $? -ne 0 ]; then
        echo "Error check failed - abort deployment"
        kubectl apply -f $MODE/canary_data_service_0.yaml
        kubectl delete deploy $DEPLOY_NAME
        exit
    else
        echo "Error check passed"
    fi

    ./performance-check.sh $VERSION
    if [ $? -ne 0 ]; then
        echo "Performance check failed - abort deployment"
        kubectl apply -f $MODE/canary_data_service_0.yaml
        kubectl delete deploy $DEPLOY_NAME
        exit
    else
        echo "Performance check passed"
    fi

    ./requests-check.sh $VERSION
    if [ $? -ne 0 ]; then
        echo "Requests check failed - abort deployment"
        kubectl apply -f $MODE/canary_data_service_0.yaml
        kubectl delete deploy $DEPLOY_NAME
        exit
    else
        echo "Requests check passed"
    fi

}

echo "Initialize!"

echo "Deploy v0.1"
kubectl apply -f $MODE/canary_data_service_0.yaml

echo "Observe v1 deployment..."
observe 30

echo "Send the v0.2 canary down the coalmine!"
kubectl apply -f $MODE/canary_data_service_3.yaml

echo "Observe the canary..."
observe 30

echo "Check for failure conditions..."
deploy_check

echo "Canary is still chirping - send a few miners!"
kubectl apply -f $MODE/canary_data_service_20.yaml

echo "Observe the miners..."
observe 20

echo "Check for failure conditions..."
deploy_check

echo "Miners are still breathing - send some more!"
kubectl apply -f $MODE/canary_data_service_50.yaml

echo "Observe the additional miners..."
observe 30

echo "Check for failure conditions..."
deploy_check

echo "All looks good - send 'em all!"
kubectl apply -f $MODE/canary_data_service_100.yaml

echo "Observe the glory..."
observe 30

echo "Clean up the old v0.1 deployment"
kubectl delete deploy meshlab-canary-data-svc-v0-1

