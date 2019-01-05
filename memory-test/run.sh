#!/bin/bash

set -ex

# REMOVE last deploy
container_id=`docker ps | grep memory-test | awk -F " " '{print $1}'`
if [ -n "$container_id" ]; then
  docker stop $container_id
  docker rm $container_id
fi

# REMOVE images
docker-compose down
docker rmi memory-test

# START new local deploy
docker build --rm -t memory-test .
docker-compose up -d
container_id=`docker ps | grep  memory-test | awk -F " " '{print $1}'`
echo "container_id: $container_id"
docker exec -it $container_id bash

docker run -d memory-test

