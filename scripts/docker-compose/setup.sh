#!/bin/sh

# Create an external network for docker
docker network create --driver bridge go-project-layout-network || true
