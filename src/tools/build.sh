#!/bin/bash

echo "Building all binaries..."
GOOS=linux GOARCH=386 go build shell-reverse-tcp.go
GOOS=windows GOARCH=386 go build shell-reverse-tcp.go
GOOS=linux GOARCH=386 go build shell-reverse-tcp-tls.go
GOOS=windows GOARCH=386 go build shell-reverse-tcp-tls.go
GOOS=linux GOARCH=386 go build shell-reverse-udp.go
GOOS=windows GOARCH=386 go build shell-reverse-udp.go
GOOS=linux GOARCH=386 go build smuggler.go
GOOS=windows GOARCH=386 go build smuggler.go
GOOS=linux GOARCH=386 go build smuggler-tls.go
GOOS=windows GOARCH=386 go build smuggler-tls.go
