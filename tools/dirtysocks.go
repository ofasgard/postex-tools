package main
//Opens a local SOCKS proxy on the port of your choice.

import "fmt"
import "os"
import "strconv"
import "github.com/ofasgard/postex-tools/postex"

func main() {
	//input handling
	if len(os.Args) < 2 {
		fmt.Println("SYNTAX:", os.Args[0], "[port]")
		return
	}
	port,err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("SYNTAX:", os.Args[0], "[port]")
		return
	}
	//let's do it!
	fmt.Println("Opening a SOCKS proxy on port " + os.Args[1])
	//this is temporary
	postex.StartSOCKSProxy(port)
}
