#!/bin/bash
# wait-for-posgres.sh

set -e

cmd="$@"

until psql -h "$DB_HOST" -p $DB_PORT -U $DB_USER -w -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd