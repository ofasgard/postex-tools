package main
//Send or receive a file by connecting to a remote TCP port.

import "postex"
import "fmt"
import "os"
import "strconv"
import "flag"

func main() {
	//input handling
	flag.Usage = usage
	var tls_ptr = flag.Bool("tls", false, "")
	flag.Parse()
	tls := *tls_ptr
	//Check we have enough positional arguments.
	if flag.NArg() != 4 {
		usage()
		return
	}
	keyword := flag.Arg(0)
	filepath := flag.Arg(1)
	host := flag.Arg(2)
	port,err := strconv.Atoi(flag.Arg(3))
	if err != nil {
		usage()
		return
	}
	//let's do it!
	switch keyword {
		case "get":
			if tls {
				err = postex.RecvFileTLS(filepath, host, port, true)
			} else {
				err = postex.RecvFile(filepath, host, port)
			}
			if err != nil {
				fmt.Println("[?]", err)
				return
			}
		case "send":
			if tls {
				err = postex.SendFileTLS(filepath, host, port, true)
			} else {
				err = postex.SendFile(filepath, host, port)
			}
			if err != nil {
				fmt.Println("[?]", err)
				return
			}
		default:
			fmt.Println("SYNTAX:", os.Args[0], "get|send [file] [target] [port]")
			return
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "SYNTAX: %s get|send [file] [target] [port]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Optional Flags:\n")
	fmt.Fprintf(os.Stderr, "\t--tls\tSet this flag to send or receive over TLS.\n")
	fmt.Fprintf(os.Stderr, "\nExample: %s --tls send file.txt 192.168.1.27 8080\n", os.Args[0])
}
