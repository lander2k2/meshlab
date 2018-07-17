#!/bin/bash

TARGET_VERSION=$1

RESP=$(curl -g 'http://localhost:9090/api/v1/query?query=istio_request_duration_bucket{destination_service="meshlab-canary-data-svc.default.svc.cluster.local"}' | jq '.data.result[] | select(.metric.destination_version == "'"$TARGET_VERSION"'") | select(.metric.le == "1") | select(.metric.source_service == "meshlab-ui.default.svc.cluster.local") | select(.metric.response_code == "200") | .value[1]')

if [ "$RESP" == "\"0\"" ]; then
    exit 1
fi

exit 0

