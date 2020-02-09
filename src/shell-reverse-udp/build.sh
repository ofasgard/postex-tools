#!/bin/bash

echo "Building shell-reverse-udp binaries..."
GOOS=linux GOARCH=386 go build shell-reverse-udp.go
GOOS=windows GOARCH=386 go build shell-reverse-udp.go

