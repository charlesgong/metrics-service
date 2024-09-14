#!/usr/bin/env bash

echo "stop metrics servers"
docker stop grafana pushgateway prometheus

echo "remove metrics servers"
docker rm grafana pushgateway prometheus