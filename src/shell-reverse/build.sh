#!/bin/bash

echo "Building shell-reverse binaries..."
GOOS=linux GOARCH=386 go build shell-reverse.go
GOOS=windows GOARCH=386 go build shell-reverse.go
