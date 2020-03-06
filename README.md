# postex-tools

A set of libraries and accompanying tools for post-exploitation, written in Golang. Intended to be a simple and straightforward toolbox for post-exploitation binaries that you can easily cross compile. For use on engagements where a standard Meterpreter payload won't cut it. 

This project is split into two parts:

- The `postex` package contains various functions that should be useful for general post-exploitation.
- The other folders in this repository contain sample Go programs that use the postex package to do things like open reverse shells or exfiltrate data.

## Tool List

Currently, this project includes the following functional tools, built using the postex package:

- `shell-reverse` is an ncat-style reverse shell. Supports shells sent over TCP, UDP and TLS (compatible with the '-ssl' option for ncat). Also supports an HTTPS shell that sends base-64 encoded input and output via GET and POST requests.
- `smuggler` is a tool for sending or receiving files by connecting to a remote host over TCP. There is a cleartext and TLS version (compatible with the '-ssl' option for ncat).
- `dirtysocks` is a portable SOCKS proxy that can be dropped onto a server and used for pivoting via SSH port forwarding.
- `shellcode` is a simple tool for executing shellcode hex strings on Windows or Linux. Try using it to execute a meterpreter payload!

Each of the tool folders includes a `build.sh` script which will automatically compile 32-bit UNIX and Windows binaries.
