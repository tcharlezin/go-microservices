#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o authApp ./cmd/api
docker build --platform linux/x86_64 -f authentication-service.dockerfile -t tcharlezin/authentication-service:1.0.0 .
docker push tcharlezin/authentication-service:1.0.0
rm authApp