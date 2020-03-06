#!/bin/bash

export GOPATH=`pwd`
export GOBIN=`pwd`/bin

echo "Building all binaries..."
GOOS=linux GOARCH=386 go build -o bin/dirtysocks src/dirtysocks/dirtysocks.go
GOOS=windows GOARCH=386 go build -o bin/dirtysocks.exe src/dirtysocks/dirtysocks.go 
GOOS=linux GOARCH=386 CGO_ENABLED=1 go build -o bin/shellcode src/shellcode/shellcode-linux.go
GOOS=windows GOARCH=386 go build -o bin/shellcode.exe src/shellcode/shellcode-windows.go
GOOS=linux GOARCH=386 go build -o bin/shell-reverse src/shell-reverse/shell-reverse.go
GOOS=windows GOARCH=386 go build -o bin/shell-reverse.exe src/shell-reverse/shell-reverse.go
GOOS=linux GOARCH=386 go build -o bin/smuggler src/smuggler/smuggler.go
GOOS=windows GOARCH=386 go build -o bin/smuggler.exe src/smuggler/smuggler.go
GOOS=linux GOARCH=386 go build -o bin/smuggler-tls src/smuggler/smuggler-tls.go
GOOS=windows GOARCH=386 go build -o bin/smuggler-tls.exe src/smuggler/smuggler-tls.go
