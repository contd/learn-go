#!/usr/loca/env bash

# Clean up
sed -i.bak "s/\"type\":/\"category\":/" links.json
sed -i.bak "s/\"done\": true/\"done\": \"1\"/" links.json
sed -i.bak "s/\"done\": false/\"done\": \"0\"/" links.json

# MongoDB Import from JSON
mongoimport -d saved -c links --jsonArray links.json

# MongoDB Export to CSV
mongoexport -d saved -c links -f url,category,created_on,done --csv > links.csv

# Create MYSQL table
mysql saved < links_create_mysql.sql

# Import to MYSQL
mysqlimport --fields-optionally-enclosed-by=\" --fields-terminated-by=, --ignore-lines=1 --columns=url,category,created_on,done --local saved links.csv

# Dump MYSQL to csv
mysql saved -e "SELECT * FROM links;" | sed "s/'/\'/;s/\t/\",\"/g;s/^/\"/;s/$/\"/;s/\n//g" > _links_mysql.csv

# SQLITE create DB and links table
cat links_create_sqlite.sql | sqlite3 saved.sqlite

# Remove first line of links_mysql.csv
len=$(cat _links_mysql.csv | wc -l)
tail -$(($len-1)) _links_mysql.csv > links_mysql.csv

# SQLITE import
sqlite3 saved.sqlite '.mode csv' '.import links_mysql.csv links'
