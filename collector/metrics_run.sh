#!/usr/bin/env bash

./metrics-stop.sh 

docker-compose --env-file config.env up -d 