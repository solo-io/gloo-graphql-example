#!/bin/bash

# build-todo-app.sh ${VERSION}

VERSION=$1
IMAGE_REPO=$2
echo $VERSION
if [[ $VERSION ]]
then
    docker build -f ./code/todo-app/Dockerfile -t ${IMAGE_REPO}:${VERSION}-amd64 --platform linux/amd64 . 
    docker build -f ./code/todo-app/Dockerfile -t ${IMAGE_REPO}:${VERSION}-arm64 --platform linux/arm64 . 
    echo "created the following images"
    echo "$IMAGE_REPO:$VERSION-arm64"
    echo "$IMAGE_REPO:$VERSION-amd64"
else
    echo "please provide a VERSION"
fi