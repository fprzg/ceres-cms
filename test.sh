#! /bin/bash

#curl -X POST -H "Content-Type: application/json" -d '{ "content": { "key": "value" } }'  \
	#http://localhost:4000/v1/api/res/proj/asfd

curl -X GET http://localhost:4000/v1/api/res/proj/en/index.yaml
