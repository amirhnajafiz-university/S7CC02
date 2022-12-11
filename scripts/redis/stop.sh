#!/bin/bash

# stop docker container
docker stop cryptometer-redis-cluster

# delete redis container image
docker image rm cryptometer-redis-cluster