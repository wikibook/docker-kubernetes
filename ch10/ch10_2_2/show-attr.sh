#!/bin/bash

ATTR=$1
if [ "$ATTR" = "" ]; then
  echo "required attribute name argument" 1>&2
  exit 1
fi

echo '{
  "id": 100,
  "username": "gihyo",
  "comment": "I like Alpine Linux"
}' | jq -r ".$ATTR"

