#!/bin/bash

IMAGE_NAME="my-mysql"

CONTAINER_NAME="my-mysql-db"

VOLUME_NAME="mysql-data"

if ! command -v podman &> /dev/null; then
  echo "Error: Podman is not installed."
  exit 1
fi

podman build -t $IMAGE_NAME .

podman run -d \
  --name $CONTAINER_NAME \
  -v $VOLUME_NAME:/Users/arthuraguiar/ContainerVolumes/mysql \
  -p 3306:3306 \
  $IMAGE_NAME
