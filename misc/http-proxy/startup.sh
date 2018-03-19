#!/bin/sh
set -e

# Write nginx config
envsubst '\${HTTP_PROXY_TARGET}' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf

# Trap SIGTERM and forward it to nginx
_term() {
  kill -TERM $child
}
trap _term SIGTERM

# Start nginx and wait for it to terminate
nginx -g 'daemon off;' &
child=$!
wait $child
