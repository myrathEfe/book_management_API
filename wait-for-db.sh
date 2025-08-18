#!/bin/sh
echo "Waiting for database to be ready..."
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  sleep 2
done
echo "Database is ready, starting app..."
exec "$@"