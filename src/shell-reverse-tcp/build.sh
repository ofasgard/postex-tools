#!/bin/bash

echo "Building shell-reverse-tcp binaries..."
GOOS=linux GOARCH=386 go build shell-reverse-tcp.go
GOOS=windows GOARCH=386 go build shell-reverse-tcp.go
GOOS=linux GOARCH=386 go build shell-reverse-tcp-tls.go
GOOS=windows GOARCH=386 go build shell-reverse-tcp-tls.go
