#!/bin/sh
# This is a test script to build the go app, locally

CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o bin/redirect
