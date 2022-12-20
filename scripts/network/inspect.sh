#!/bin/bash

# inspect container network
docker inspect $1 -f "{{json .NetworkSettings.Networks }}"
