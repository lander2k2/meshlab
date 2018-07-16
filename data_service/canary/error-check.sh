#!/bin/bash

RESP=$(curl -g 'http://localhost:9090/api/v1/query?query=istio_request_count{destination_service="meshlab-canary-data-svc.default.svc.cluster.local"}')
INDEX="0"
COMPLETE="false"

R200=0
R500=0

TARGET_VERSION=$1

while [ "$COMPLETE" == "false" ]; do
    RESULT=$(echo $RESP | jq ".data.result[$INDEX]")
    if [ "$RESULT" == "null" ]; then
        COMPLETE="true"
    fi

    CODE=$(echo $RESULT | jq ".metric.response_code")
    COUNT=$(echo $RESULT | jq ".value[1]")
    VERSION=$(echo $RESULT | jq ".metric.destination_version")

    echo "$VERSION: $COUNT x $CODE response code"

    if (( "${CODE//\"}" == 500 )) \
        && (( "${COUNT//\"}" > 10 )) \
        && [ "${VERSION//\"}" == "${TARGET_VERSION}" ]; then
            exit 1
    fi

    INDEX=$((INDEX+1))

done

exit 0

