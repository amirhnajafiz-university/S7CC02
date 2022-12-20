#!/bin/bash

# running redis on docker
docker run --network=cryptonet -v redisvolume:/data --name cryptometer-redis-cluster -d redis:latest
