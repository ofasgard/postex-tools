# postex-tools

A set of libraries and accompanying tools for post-exploitation, written in Golang. Intended to be a simple-straightforward way to cross-compile post exploitation binaries for use on engagements where a standard Meterpreter payload won't cut it. This project is split into two parts:

- The postex package contains various functions that should be useful for general post-exploitation.
- The tools folder contains Go programs that use the postex package to do things like open reverse shells or exfiltrate data.

## Building

If you want to build the tools in the tools folder, it should be enough to clone this repository, set it as your GOPATH, and run the build.sh file in the tools folder.
