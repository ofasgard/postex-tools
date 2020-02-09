#!/bin/bash

echo "Building smuggler binaries..."
GOOS=linux GOARCH=386 go build smuggler.go
GOOS=windows GOARCH=386 go build smuggler.go
GOOS=linux GOARCH=386 go build smuggler-tls.go
GOOS=windows GOARCH=386 go build smuggler-tls.go
