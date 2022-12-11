#!/bin/bash

# running redis on docker
docker run -v redisvolume:/data --name cryptometer-redis-cluster -d redis