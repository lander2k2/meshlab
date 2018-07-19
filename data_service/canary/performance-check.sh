#!/bin/bash

TARGET_VERSION=$1

RESP=$(curl -g 'http://localhost:9090/api/v1/query?query=istio_request_duration_bucket{destination_service="meshlab-canary-data-svc.default.svc.cluster.local"}')

REQ_DURATION_BUCKET=$(echo $RESP | jq '.data.result[] | select(.metric.destination_version == "'"$TARGET_VERSION"'") | select(.metric.le == "1") | select(.metric.source_service == "meshlab-ui.default.svc.cluster.local") | select(.metric.response_code == "200")')

REQ_COUNT=$(echo $RESP | jq '.data.result[] | select(.metric.destination_version == "'"$TARGET_VERSION"'") | select(.metric.le == "1") | select(.metric.source_service == "meshlab-ui.default.svc.cluster.local") | select(.metric.response_code == "200") | .value[1]')

echo "Request Duration Bucket: $REQ_DURATION_BUCKET"

if [ "$REQ_COUNT" == "\"0\"" ]; then
    exit 1
fi

exit 0

