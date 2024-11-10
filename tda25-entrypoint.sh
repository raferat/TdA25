#!/bin/sh

/usr/local/bin/docker-entrypoint.sh postgres &

sleep 1m

BASEDIR=$(dirname "$0")

cd $BASEDIR && ./tdaserver


