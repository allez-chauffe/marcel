#!/bin/bash
# Script used to build docker images on CircleCI
# $1 : image name (without version)
set -e
set -o pipefail
# Any subsequent(*) commands which fail will cause the shell script to exit immediately

IMAGE_NAME=${1}

PUSH=false
if [[ "${CIRCLE_BRANCH}" == "master" || "${CIRCLE_TAG}" != "" ]] ; then
  PUSH=true
fi

IMAGE_VERSION=dev
if [[ "${CIRCLE_TAG}" != "" ]] ; then
  IMAGE_VERSION="$CIRCLE_TAG"
  PUSH=true
fi

echo "IMAGE_NAME    = ${IMAGE_NAME}"
echo "IMAGE_VERSION = ${IMAGE_VERSION}"
echo "PUSH          = ${PUSH}"

# Build docker image.
docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .

# Push image to docker hub.
if [ "${PUSH}" == "true" ] ; then
  echo "authenticate to google cloud"
  echo ${GCLOUD_SERVICE_KEY} | base64 --decode --ignore-garbage > ${HOME}/gcloud-service-key.json
  gcloud auth activate-service-account --key-file=${HOME}/gcloud-service-key.json
  gcloud config set project ${GCLOUD_PROJECT_ID}
  gcloud auth configure-docker eu.gcr.io

  echo "push docker image"
  docker push ${IMAGE_NAME}:${IMAGE_VERSION}
fi
