#!/bin/bash
set -e

if [ -z $POSTGRES_USER ]; then
    dropdb --if-exists tda25
    dropuser --if-exists goserver
    POSTGRES_USER="$USER"
fi

BASEDIR=$(dirname "$0")

createdb --username "$POSTGRES_USER" tda25
psql --username "$POSTGRES_USER" -f "$BASEDIR/schemas/createuser.sql" tda25
psql --username "$POSTGRES_USER" -f "$BASEDIR/schemas/gametable.sql" tda25

