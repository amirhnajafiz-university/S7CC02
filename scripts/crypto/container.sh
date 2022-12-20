#!/bin/bash

# running crypto meter on docker
docker run --network=cryptonet --name cryptometer -d -p 6030:6030 --mount type=bind,source="$(pwd)"/config.yml,target=/app/config.yml amirhossein21/cryptometer:v0.1
