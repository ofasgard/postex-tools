#!/bin/bash

export GOARCH=amd64

echo "Building all binaries for '$GOARCH'..."

GOOS=linux go build -o bin/linux/dirtysocks tools/dirtysocks.go
GOOS=windows go build -o bin/windows/dirtysocks.exe tools/dirtysocks.go

GOOS=linux CGO_ENABLED=1 go build -o bin/linux/shellcode tools/shellcode-linux.go
GOOS=windows go build -o bin/windows/shellcode.exe tools/shellcode-windows.go

GOOS=linux CGO_ENABLED=1 go build -o bin/linux/shellcode-inject tools/shellcode-inject-linux.go
GOOS=windows go build -o bin/windows/shellcode-inject.exe tools/shellcode-inject-windows.go

GOOS=linux go build -o bin/linux/shell-reverse tools/shell-reverse.go
GOOS=windows go build -o bin/windows/shell-reverse.exe tools/shell-reverse.go

GOOS=linux go build -o bin/linux/smuggler tools/smuggler.go
GOOS=windows go build -o bin/windows/smuggler.exe tools/smuggler.go

GOOS=linux go build -o bin/linux/xortool tools/xortool.go
GOOS=windows go build -o bin/windows/xortool.exe tools/xortool.go
