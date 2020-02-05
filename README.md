# postex-tools

A set of libraries and accompanying tools for post-exploitation, written in Golang. Intended to be a simple and straightforward toolbox for post-exploitation binaries that you can easily cross compile. For use on engagements where a standard Meterpreter payload won't cut it. This project is split into two parts:

- The postex package contains various functions that should be useful for general post-exploitation.
- The tools folder contains sample Go programs that use the postex package to do things like open reverse shells or exfiltrate data.

TODO:

- Add support for encrypted shells.
- Add support for DNS tunneling.
- Add support for tunneling over HTTPS.

## Building

If you want to build the tools in the `tools/` folder, it should be enough to clone this repository, set it as your GOPATH, and run the build script in the `tools/` folder.
