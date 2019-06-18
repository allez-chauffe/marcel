#!/bin/bash
# Script used to install marcel on GCP virtual machines
# $1 : version to install
set -e
set -o pipefail
# Any subsequent(*) commands which fail will cause the shell script to exit immediately

if [ ! -z "${1}" ] ; then
  VERSION=${1}
elif [[ "${CIRCLE_TAG}" != "" ]] ; then
  VERSION=${CIRCLE_TAG}
else
  VERSION=dev
fi

if [ "${CI}" == true ] ; then
  echo "authenticate to google cloud"
  echo ${GCLOUD_SERVICE_KEY} | base64 --decode --ignore-garbage > ${HOME}/gcloud-service-key.json
  gcloud auth activate-service-account --key-file=${HOME}/gcloud-service-key.json
  gcloud config set project ${GCLOUD_PROJECT_ID}

  echo "copy ssh key"
  mkdir -p ${HOME}/.ssh
  chmod 700 ${HOME}/.ssh
  echo ${GCLOUD_SSH_KEY} | base64 --decode --ignore-garbage > ${HOME}/.ssh/google_compute_engine
  chmod 600 ${HOME}/.ssh/google_compute_engine
  echo ${GCLOUD_SSH_PUB_KEY} | base64 --decode --ignore-garbage > ${HOME}/.ssh/google_compute_engine.pub
  chmod 644 ${HOME}/.ssh/google_compute_engine.pub
fi

# Use gcloud command to list servers using this version
INSTANCES=$(gcloud compute instances list --filter="labels.version = ${VERSION}" --uri)

# Do install on each instance found
PATTERN="https://www.googleapis.com/compute/v1/projects/zen-dashboard/zones/\([^/]*\)/instances/\([^/]*\)"
for url in ${INSTANCES} ; do
  zone=$(echo "${url}" | sed -e "s,${PATTERN},\1,")
  instance=$(echo "${url}" | sed -e "s,${PATTERN},\2,g")
  echo "install version ${VERSION} on instance ${zone}/${instance}"
  gcloud compute scp /home/circleci/project/plugins.tar.gz "ubuntu@${instance}:~/" --strict-host-key-checking=no --zone="${zone}"
  gcloud compute ssh "ubuntu@${instance}" --strict-host-key-checking=no --zone="${zone}" -- "./MARCEL/deploy.sh"
done
