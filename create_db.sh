#!/bin/bash
set -e

IFS=',' DB_DATA=( ${EXTRA_DB} )

for DB in ${DB_DATA[*]}; do

IFS=':' DB_SPECS=( ${DB} )
DB_NAME=${DB_SPECS[0]}
DB_USER=${DB_SPECS[1]}

echo "Creating database: ${DB_NAME} owned by ${DB_USER}"

createuser -U ${POSTGRES_USER} ${DB_USER};
createdb -U ${POSTGRES_USER} ${DB_NAME} -O ${DB_USER};

done
