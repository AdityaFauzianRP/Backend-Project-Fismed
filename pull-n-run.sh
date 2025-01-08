#!/bin/bash

set -e

git pull

git checkout main


docker pull boyangyang/fismed-be-staging:latest

# Stop and remove the fismed-be-staging-dev container if it exists
# if [ "$(docker ps -q -f name=fismed-be-staging-dev)" ]; then
#     docker stop fismed-be-staging-dev
#     docker rm fismed-be-staging-dev
# fi

# Tag the pulled image as fismed-be-staging:latest
docker tag boyangyang/fismed-be-staging:latest fismed-be-staging:latest

sleep 1

docker compose up -d
