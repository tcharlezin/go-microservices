#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o loggerServiceApp ./cmd/api
docker build --platform linux/x86_64 -f logger-service.dockerfile -t tcharlezin/logger-service:1.0.1 .
docker push tcharlezin/logger-service:1.0.1
rm loggerServiceApp