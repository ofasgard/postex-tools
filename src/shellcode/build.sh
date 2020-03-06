#!/bin/bash

echo "Building shellcode binaries..."
GOOS=linux GOARCH=386 CGO_ENABLED=1 go build shellcode-linux.go
GOOS=windows GOARCH=386 go build shellcode-windows.go
