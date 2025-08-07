#!/bin/bash

set -e

dbname="booklib"
user="booklib"
password="booklib"
host="${PGHOST:-localhost}"
port="${PGPORT:-5656}"

echo "🛠 Creating database '$dbname'..."

PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -tc "SELECT 1 FROM pg_database WHERE datname = '$dbname'" | grep -q 1 || \
PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -c "CREATE DATABASE $dbname"

echo "✅ Database '$dbname' ready."
