#! /bin/bash

BODY=

curl -X POST -H "Content-Type: application/json" -d '{ "content": { "key": "value" } }'  \
	http://localhost:4000/v1/api/res/proj/asfd
