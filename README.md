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
- `shellcode` is a simple tool for executing shellcode on Windows or Linux - provide it with a hex string or a path to a binary file containing shellcode. Try using it to execute a meterpreter payload!
- `shellcode-inject` is identical to the previous tool, but allows you to inject shellcode into an existing process by providing a PID.

## Building

This project doesn't have any external dependencies besides Go itself. To build it, just do:

```shell
$ git clone https://github.com/ofasgard/postex-tools
$ cd postex-tools
$ ./build.sh
```

If you have trouble building or using any of the tools, the following notes may be helpful:

- You may need to install `gcc-multilib` or the equivalent for cross-platform compilation to work.
- In order to use the shellcode loader, you'll need to set the GOARCH variable to the correct architecture - `i386` for 32-bit shellcode, and `amd64` for 64-bit shellcode.
