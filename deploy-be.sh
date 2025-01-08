#!/bin/bash

set -e 

docker build -t fismed-be-staging:latest .

docker compose down -v

docker compose up -d