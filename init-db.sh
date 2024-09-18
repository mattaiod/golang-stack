#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  DO
  \$do\$
  BEGIN
    IF NOT EXISTS (
      SELECT
      FROM   pg_catalog.pg_database
      WHERE  datname = '$POSTGRES_DB') THEN
      PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE $POSTGRES_DB');
    END IF;
  END
  \$do\$;
EOSQL