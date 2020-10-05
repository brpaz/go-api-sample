#!/usr/bin/env bash
set -e

echo "Running Entrypoint"

# waits for database to be up
wait-for-it -h ${DB_HOST} -p ${DB_PORT} -t 30 -q

## Did we found IP address? Use exit status of the grep command ##
if [ $? -ne 0 ]
then
  echo "Failed to ping database host"
  exit -1
fi

export PGPASSWORD=${DB_PASSWORD}

migrate -path=migrations -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable up

exec "$@"
