#!/bin/bash

set -e

# Configurable DB connection
dbname="booklib"
user="booklib"
password="booklib"
host="${PGHOST:-localhost}"
port="${PGPORT:-5656}"

echo "⏬ Running DOWN migrations on database '$dbname'..."

for file in $(ls down/*.sql | sort -r); do
  echo "⏹  Reverting $file"
  PGPASSWORD=$password psql -U $user -d "$dbname" -h $host -p $port -f "$file"
done

echo "✅ Migrations DOWN completed."
