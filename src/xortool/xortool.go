package main
//Apply XOR encryption to a hex string or file.

import "postex"
import "fmt"
import "os"
import "encoding/hex"

func main() {
	if len(os.Args) < 3 { 
		fmt.Println("USAGE: " + os.Args[0] + " <hex string> <hex key>")
		return
	}
	payload_str := os.Args[1]
	key_str := os.Args[2]
	//treat both arguments as a hex string of bytes	
	payload,err := hex.DecodeString(payload_str)
	key,err2 := hex.DecodeString(key_str)
	if (err == nil) && (err2 == nil) {
		result := postex.Xorify(payload, key)
		result_str := hex.EncodeToString(result)
		fmt.Println(result_str)
		return
	}
	//return an error
	fmt.Println("USAGE: " + os.Args[0] + " <hex string> <hex key>")	
}
