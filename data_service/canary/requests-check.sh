#!/bin/bash

TARGET_VERSION=$1

RESP=$(curl -g 'http://localhost:9090/api/v1/query?query=istio_request_count{destination_service="meshlab-canary-data-svc.default.svc.cluster.local"}' | jq '.data.result[] | select(.metric.destination_version == "'"$TARGET_VERSION"'") | select(.metric.source_service == "meshlab-ui.default.svc.cluster.local") | select(.metric.response_code == "200") | .value[1]' | tr -d '"')

echo "Client requests query response: $RESP"

# zero requests - no result at all for the query
if [ "$RESP" == "" ]; then
    exit 1
fi

# between 1 and 9 - a result with a value less then 10
if [ "$RESP" -lt "10" ]; then
    exit 1
fi

exit 0

