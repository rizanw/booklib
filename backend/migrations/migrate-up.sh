#!/bin/bash

set -e

# Configurable DB connection
dbname="booklib"
user="booklib"
password="booklib"
host="${PGHOST:-localhost}"
port="${PGPORT:-5656}"

echo "🔼 Running UP migrations on database '$dbname'..."

for file in $(ls up/*.sql | sort); do
  echo "▶️  Applying $file"
  PGPASSWORD=$password psql -U $user -d "$dbname" -h $host -p $port -f "$file"
done

echo "✅ Migrations UP completed."
