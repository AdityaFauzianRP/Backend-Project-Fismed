#!/bin/bash

set -e

git pull

sleep 1

git checkout main

sleep 1

docker buildx build --platform linux/amd64 --progress=plain -t boyangyang/fismed-be:latest .

sleep 1

docker push boyangyang/fismed-be:latest
