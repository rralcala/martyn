#!/bin/bash

set -x
curl http://localhost:8080/transactions
curl http://localhost:8080/transactions/3
curl http://localhost:8080/transactions/4
curl http://localhost:8080/transactions --include --header "Content-Type: application/json"  --request "POST" \
    --data '{"date": "2020-12-09T16:09:53+00:00","Description": "Initial deposit","amount": 1528000, "cost_center": 1, "account": 3, "provider": 2}'
#curl http://localhost:8080/transactions/4
