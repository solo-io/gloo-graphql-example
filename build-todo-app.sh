#!/bin/bash

# VERSION=0.0.3 IMAGE_REPO=my_repo build-todo-app.sh 

echo $VERSION $IMAGE_REPO
if [[ $VERSION ]] && [[ $IMAGE_REPO ]]
then
    docker build -f ./code/todo-app/Dockerfile -t ${IMAGE_REPO}:${VERSION}-amd64 --platform linux/amd64 . 
    docker build -f ./code/todo-app/Dockerfile -t ${IMAGE_REPO}:${VERSION}-arm64 --platform linux/arm64 . 
    echo "created the following images"
    echo "$IMAGE_REPO:$VERSION-arm64"
    echo "$IMAGE_REPO:$VERSION-amd64"
else
    echo "please provide a VERSION and a IMAGE_REPO as environment variables"
fi