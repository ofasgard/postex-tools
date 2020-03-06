package main
//Execute shellcode locally. If it's 32-bit shellcode, compile this as a 32-bit binary.

import "postex"
import "fmt"
import "os"
import "encoding/hex"

func main() {
	if len(os.Args) < 2 { 
		fmt.Println("USAGE: " + os.Args[0] + " <shellcode>")
		return
	}
	sc_str := os.Args[1]
	sc,err := hex.DecodeString(sc_str)
	if err != nil {
		fmt.Println("Couldn't decode the shellcode. Give me a hex string!")
		return
	}
	postex.ShellcodeWindows(sc)
}
