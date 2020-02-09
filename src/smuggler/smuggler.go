package main
//Send or receive a file by connecting to a remote TCP port.

import "postex"
import "fmt"
import "os"
import "strconv"

func main() {
	//input handling
	if len(os.Args) < 5 {
		fmt.Println("SYNTAX:", os.Args[0], "get|send [file] [target] [port]")
		return
	}
	keyword := os.Args[1]
	filepath := os.Args[2]
	host := os.Args[3]
	port,err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("SYNTAX:", os.Args[0], "get|send [file] [target] [port]")
		return
	}
	//let's do it!
	switch keyword {
		case "get":
			err := postex.RecvFile(filepath, host, port)
			if err != nil {
				fmt.Println("[?]", err)
				return
			}
		case "send":
			err := postex.SendFile(filepath, host, port)
			if err != nil {
				fmt.Println("[?]", err)
				return
			}
		default:
			fmt.Println("SYNTAX:", os.Args[0], "get|send [file] [target] [port]")
			return
	}
}
