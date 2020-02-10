#!/bin/bash

echo "Building shell-reverse-tcp binaries..."
GOOS=linux GOARCH=386 go build shell-reverse-https.go
GOOS=windows GOARCH=386 go build shell-reverse-https.go
