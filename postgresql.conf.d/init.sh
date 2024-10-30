#!/bin/bash
set -e

psql << EOF
  DROP DATABASE server;
EOF

createdb server
psql -f initsql/createuser.sql server
psql -f initsql/gametable.sql server

