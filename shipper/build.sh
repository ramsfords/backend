#!/bin/bash

GO111MODULE=on
CGO_ENABLED=0


go mod vendor

go build -ldflags="-w -s" -o main  
