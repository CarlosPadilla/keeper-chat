#!/usr/bin/env bash

DATABASE_STRING="postgres://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

echo "Running database migrations: (${DATABASE_STRING})"
migrate -database ${DATABASE_STRING} -path fixtures up

./run.sh