#!/bin/bash

export GOPATH=`pwd`
export GOBIN=`pwd`/bin

export GOARCH=amd64

echo "Building all binaries for '$GOARCH'..."

GOOS=linux go build -o bin/linux/dirtysocks src/dirtysocks/dirtysocks.go
GOOS=windows go build -o bin/windows/dirtysocks.exe src/dirtysocks/dirtysocks.go 

GOOS=linux CGO_ENABLED=1 go build -o bin/linux/shellcode src/shellcode/shellcode-linux.go
GOOS=windows go build -o bin/windows/shellcode.exe src/shellcode/shellcode-windows.go

GOOS=linux CGO_ENABLED=1 go build -o bin/linux/shellcode-inject src/shellcode-inject/shellcode-inject-linux.go
GOOS=windows go build -o bin/windows/shellcode-inject.exe src/shellcode-inject/shellcode-inject-windows.go

GOOS=linux go build -o bin/linux/shell-reverse src/shell-reverse/shell-reverse.go
GOOS=windows go build -o bin/windows/shell-reverse.exe src/shell-reverse/shell-reverse.go

GOOS=linux go build -o bin/linux/smuggler src/smuggler/smuggler.go
GOOS=windows go build -o bin/windows/smuggler.exe src/smuggler/smuggler.go
