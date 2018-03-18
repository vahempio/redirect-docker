#!/bin/sh
# This is a script to build the go app inside the golang docker

docker run \
	--rm \
	--user "`id -u`:`id -g`" \
	-v "$PWD":/usr/src/myapp \
	-w /usr/src/myapp \
	-e CGO_ENABLED=0 \
	golang:1.10.0 go build -a --installsuffix cgo --ldflags="-s" -o bin/redirect
