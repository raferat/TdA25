#!/bin/bash
set -e

psql << EOF
  DROP DATABASE server;
EOF
psql -f initsql/createdb.sql
psql -f initsql/createuser.sql server
psql -f initsql/gametable.sql server

