#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o mailerApp ./cmd/api
docker build --platform linux/x86_64 -f mail-service.dockerfile -t tcharlezin/mail-service:1.0.0 .
docker push tcharlezin/mail-service:1.0.0
rm mailerApp