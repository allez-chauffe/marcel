#!/bin/bash
# Script used to build docker images on CircleCI
# $1 : image name (without version)
set -e
set -o pipefail
# Any subsequent(*) commands which fail will cause the shell script to exit immediately

PUSH=false
if [ "${CIRCLE_BRANCH}" == "master" ] ; then
  PUSH=true
fi

IMAGE_NAME=${1}

IMAGE_VERSION=dev
if [[ "${CIRCLE_BRANCH}" =~ ^release-[0-9]+\.[0-9]+\.[0-9]+$ ]] ; then
  IMAGE_VERSION=${CIRCLE_BRANCH:8} # remove 'release-' from branch name to get version
  PUSH=true
fi

echo "IMAGE_NAME    = ${IMAGE_NAME}"
echo "IMAGE_VERSION = ${IMAGE_VERSION}"
echo "PUSH          = ${PUSH}"

# Build docker image.
docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .

# Push image to docker hub.
if [ "${PUSH}" == "true" ] ; then
  docker login -u ${DOCKER_LOGIN} -p ${DOCKER_PASSWORD}
  docker push ${IMAGE_NAME}:${IMAGE_VERSION}
fi
