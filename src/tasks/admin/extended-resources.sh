#!/bin/bash

# Please set up the proxy
kubectl proxy

# Add
curl --header "Content-Type: application/json-patch+json" \
     --request PATCH \
     --data '[{"op": "add", "path": "/status/capacity/diguage.com~1dongle", "value": "4"}]' \
     http://localhost:8001/api/v1/nodes/node3/status


# Remove
curl --header "Content-Type: application/json-patch+json" \
     --request PATCH \
     --data '[{"op": "remove", "path": "/status/capacity/diguage.com~1dongle"}]' \
     http://localhost:8001/api/v1/nodes/node3/status
