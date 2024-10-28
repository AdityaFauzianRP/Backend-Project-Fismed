#!/bin/bash

set -e 

docker build -t fismed-be:latest .

docker compose up -d