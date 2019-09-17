#!/bin/bash

set -e;

echo "Setting config..."
envsubst < "$@" > "$@";

if [ -z $HOST ]; then
    exit
else
    echo "HOST = " $HOST
fi

echo "Starting..."
nginx -c "$@" -g "daemon off;";

