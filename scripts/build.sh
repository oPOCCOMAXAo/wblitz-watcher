#!/bin/sh

git fetch --tags origin

VERSION=`git describe --tags --exact-match || echo ""`
if [ -z "$VERSION" ]
then
    echo "No tags"
    return 1
fi

echo "Build "$VERSION

CGO_ENABLED=0 go build -o wbwatcher main.go

docker build -t wbwatcher:$VERSION .
docker image tag wbwatcher:$VERSION poccomaxa/wbwatcher:$VERSION
docker push poccomaxa/wbwatcher:$VERSION
