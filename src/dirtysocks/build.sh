#!/bin/bash

echo "Building dirtysocks binaries..."
GOOS=linux GOARCH=386 go build dirtysocks.go
GOOS=windows GOARCH=386 go build dirtysocks.go
