#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o listenerApp .
docker build --platform linux/x86_64 -f listener-service.dockerfile -t tcharlezin/listener-service:1.0.0 .
docker push tcharlezin/listener-service:1.0.0
rm listenerApp