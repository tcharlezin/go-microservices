#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o brokerApp ./cmd/api
docker build --platform linux/x86_64 -f broker-service.dockerfile -t tcharlezin/broker-service:1.0.1 .
docker push tcharlezin/broker-service:1.0.1
rm brokerApp