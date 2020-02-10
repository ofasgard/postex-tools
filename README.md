# postex-tools

A set of libraries and accompanying tools for post-exploitation, written in Golang. Intended to be a simple and straightforward toolbox for post-exploitation binaries that you can easily cross compile. For use on engagements where a standard Meterpreter payload won't cut it. 

This project is split into two parts:

- The postex package contains various functions that should be useful for general post-exploitation.
- The other folders in this repository contain sample Go programs that use the postex package to do things like open reverse shells or exfiltrate data.

## Tool List

Currently, this project includes the following functional tools, built using the postex package:

- `shell-reverse-tcp` is an ncat-style TCP reverse shell. There is a cleartext and TLS version (compatible with the `-ssl` option for ncat).
- `shell-reverse-udp` is an ncat-style UDP reverse shell.
- `shell-reverse-https` is a reverse shell that functions over HTTPS. Base64-encoded commands are retrieved with a GET request to /input, and base64-encoded output is returned with a POST request to /output.
- `smuggler` is a tool for for sending or receiving files by connecting to a remote host over TCP. There is a cleartext and TLS version (compatible with the `-ssl` option for ncat).

Each of the tool folders includes a `build.sh` script which will automatically compile 32-bit UNIX and Windows binaries.
