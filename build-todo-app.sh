#!/bin/bash

# build-todo-app.sh ${VERSION}

VERSION=$1
echo $VERSION
if [[ $VERSION ]]
then
    docker build -f ./code/todo-app/Dockerfile -t gcr.io/solo-test-236622/graphql-todo:${VERSION}-amd64 --platform linux/amd64 . 
    docker build -f ./code/todo-app/Dockerfile -t gcr.io/solo-test-236622/graphql-todo:${VERSION}-arm64 --platform linux/arm64 . 
    echo "ensure that this is the correct version that you want and then use the following commands to push the images"
    echo "docker push gcr.io/solo-test-236622/graphql-todo:$VERSION-amd64"
    echo "docker push gcr.io/solo-test-236622/graphql-todo:$VERSION-arm64"
else
    echo "please provide a VERSION"
fi