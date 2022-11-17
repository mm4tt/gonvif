#!/bin/bash

go run cmd/gonvif/main.go  -a http://88.135.162.219:10080 -p pulputron -u admin  $@
