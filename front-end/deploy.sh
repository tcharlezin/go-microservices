#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o frontEndApp ./cmd/web
docker build --platform linux/x86_64 -f frontend-service.dockerfile -t tcharlezin/front-end:1.0.1 .
docker push tcharlezin/front-end:1.0.1
rm frontEndApp