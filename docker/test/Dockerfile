# syntax = docker/dockerfile:experimental
# from a base image
FROM alpine:3.12

# update and upgrade
RUN apk-upgrade

# install dependencies
RUN apk add --no-cache \
    curl

# run a command
RUN curl https://www.google.com