#!/bin/bash

# build crypto image
docker build . -t amirhossein21/cryptometer:"$1" -f docker/build/Dockerfile