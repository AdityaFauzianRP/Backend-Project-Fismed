#!/bin/bash

set -e 

docker build -t fismed-be:latest .

docker compose down -v

docker compose up -d