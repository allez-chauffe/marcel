## Build
docker image -t zenika/http-proxy .

## Usage
docker container run -dP -e HTTP_PROXY_TARGET=https://zenika.com zenika/http-proxy
