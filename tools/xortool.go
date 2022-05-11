package main
//Apply XOR encryption to a hex string or file.

import "github.com/ofasgard/postex-tools/postex"
import "fmt"
import "os"
import "io/ioutil"
import "encoding/hex"

func main() {
	if len(os.Args) < 3 { 
		fmt.Println("USAGE: " + os.Args[0] + " <hex string or filename> <hex key>")
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
	//if that fails, treat the first argument as a filepath and the second as a hex key
	//in this mode, output is binary rather than hex
	fd,err := os.Open(payload_str)
	key,err2 = hex.DecodeString(key_str)
	if (err == nil) && (err2 == nil) {
		defer fd.Close()
		info,err := os.Stat(payload_str)
		if err == nil {
			size := info.Size()
			payload := make([]byte, size)
			payload,err = ioutil.ReadAll(fd)
			if err == nil {
				result := postex.Xorify(payload, key)
				fmt.Println(string(result))
				return
			}
		}
	}
	//return an error
	fmt.Println("USAGE: " + os.Args[0] + " <hex string or filename> <hex key>")	
}
