# syntax = docker/dockerfile:experimental
# from a base image
FROM alpine:3.12

# update and upgrade
RUN apk update

# install dependencies
RUN apk add --no-cache \
    curl

# cmd part
CMD ["tail", "-f", "/dev/null"]
