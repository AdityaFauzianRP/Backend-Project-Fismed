#!/bin/bash

set -e

docker pull boyangyang/fismed-be:latest

# Stop and remove the fismed-be-dev container if it exists
if [ "$(docker ps -q -f name=fismed-be-dev)" ]; then
    docker stop fismed-be-dev
    docker rm fismed-be-dev
fi

# Tag the pulled image as fismed-be:latest
docker tag boyangyang/fismed-be:latest fismed-be:latest

sleep 1

docker compose up -d
