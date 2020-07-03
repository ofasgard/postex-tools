#!/bin/bash

export GOPATH=`pwd`
export GOBIN=`pwd`/bin

export GOARCH=amd64

echo "Building all binaries for '$GOARCH'..."

GOOS=linux go build -o bin/dirtysocks src/dirtysocks/dirtysocks.go
GOOS=windows go build -o bin/dirtysocks.exe src/dirtysocks/dirtysocks.go 

GOOS=linux CGO_ENABLED=1 go build -o bin/shellcode src/shellcode/shellcode-linux.go
GOOS=windows go build -o bin/shellcode.exe src/shellcode/shellcode-windows.go

GOOS=linux CGO_ENABLED=1 go build -o bin/shellcode-inject src/shellcode-inject/shellcode-inject-linux.go
GOOS=windows go build -o bin/shellcode-inject.exe src/shellcode-inject/shellcode-inject-windows.go

GOOS=linux go build -o bin/shell-reverse src/shell-reverse/shell-reverse.go
GOOS=windows go build -o bin/shell-reverse.exe src/shell-reverse/shell-reverse.go

GOOS=linux go build -o bin/smuggler src/smuggler/smuggler.go
GOOS=windows go build -o bin/smuggler.exe src/smuggler/smuggler.go
