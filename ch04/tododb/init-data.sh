#!/bin/bash -e

for SQL in `ls /sql/*.sql`
do
  mysql -u root -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < $SQL > /dev/null 2>&1
done
