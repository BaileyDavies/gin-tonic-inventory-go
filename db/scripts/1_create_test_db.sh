!/bin/bash

set -e
echo "test!"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    DROP DATABASE IF EXISTS gintonictest;
    CREATE DATABASE gintonictest;
EOSQL