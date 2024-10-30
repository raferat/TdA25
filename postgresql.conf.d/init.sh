#!/bin/bash
set -e

psql << EOF
  DROP DATABASE tda25;
EOF

createdb tda25
psql -f schemas/createuser.sql tda25
psql -f schemas/gametable.sql tda25

