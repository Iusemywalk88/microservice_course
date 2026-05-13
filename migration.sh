#!/bin/bash
set -a
source .env
set +a

echo "Starting migrations..."
echo "MIGRATION_AUTH_DIR: ${MIGRATION_AUTH_DIR}"
echo "MIGRATION_CHAT_DIR: ${MIGRATION_CHAT_DIR}"

export MIGRATION_AUTH_DSN="host=postgres port=5432 dbname=$AUTH_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"
export MIGRATION_CHAT_DSN="host=postgres port=5432 dbname=$CHAT_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

sleep 2

echo "Running Auth migrations..."
goose -dir "${MIGRATION_AUTH_DIR}" postgres "${MIGRATION_AUTH_DSN}" up -v

echo "Running Chat migrations..."
goose -dir "${MIGRATION_CHAT_DIR}" postgres "${MIGRATION_CHAT_DSN}" up -v