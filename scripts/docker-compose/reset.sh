#!/bin/sh

# Remove an external network for docker
docker network rm go-project-layout-network
rm -rf ./.volume
