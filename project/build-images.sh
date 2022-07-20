#!/bin/bash

cd ../authentication-service && ./deploy.sh
cd ../broker-service && ./deploy.sh
cd ../front-end && ./deploy.sh
cd ../listener-service && ./deploy.sh
cd ../logger-service && ./deploy.sh
cd ../mail-service && ./deploy.sh