package main
//Execute shellcode locally. If it's 32-bit shellcode, compile this as a 32-bit binary.
//You MUST use the correct GOARCH for the shellcode!

import "postex"
import "fmt"
import "os"
import "io/ioutil"
import "encoding/hex"

func main() {
	if len(os.Args) < 2 { 
		fmt.Println("USAGE: " + os.Args[0] + " <shellcode>")
		return
	}
	sc_str := os.Args[1]
	//first try treating it as a hex string
	sc,err := hex.DecodeString(sc_str)
	if err == nil {
		postex.ShellcodeWindows(sc)
		return
	}
	//next try treating it as a filename
	fd,err := os.Open(sc_str)
	if err == nil {
		defer fd.Close()
		info,err := os.Stat(sc_str)
		if err == nil {
			size := info.Size()
			sc := make([]byte, size)
			sc,err = ioutil.ReadAll(fd)
			if err == nil {
				postex.ShellcodeWindows(sc)
				return
			}
		}
	}
	//return an error
	fmt.Println("Please provide a valid shellcode string or path of a binary file to execute.")
}
