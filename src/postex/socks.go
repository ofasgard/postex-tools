package postex
//Contains functions for implementing a SOCKS5 proxy server in Go.
//It's quick and dirty and doesn't fully implement SOCKS5, but it'll do the job.

import "net"
import "strconv"
import "fmt"
import "encoding/binary"

const SOCKS_VERSION = 5

/*
* StartSOCKSProxy(port int)
*
* Start a SOCKS5 proxy on the designated port; runs perpetually.
*/

func StartSOCKSProxy(port int) {
	//setup
	port_str := ":" + strconv.Itoa(port)
	server_addr,err := net.ResolveTCPAddr("tcp4", port_str)
	if err != nil {
		fmt.Println(string(err.Error()))
		return
	}
	//start the listener
	listener,err := net.ListenTCP("tcp4", server_addr)
	if err != nil {
		fmt.Println(string(err.Error()))
		return
	}
	defer listener.Close()
	//accept connections
	for {
		conn,err := listener.Accept()
		if err == nil {
			//use a goroutine to concurrently handle connections
			go func() {
				err := handleSOCKS(conn)
				if err != nil {
					fmt.Println(string(err.Error()))
				}
			}()
		} else {
			fmt.Println(string(err.Error()))
		}
	}
}

/*
* handleSOCKS (conn net.Conn) error
*
* Handle a single connection request from a SOCKS5 client.
* Uses goroutines to concurrently send and receive data between the client and the remote host.
*/

func handleSOCKS(conn net.Conn) error {
	//perform the initial greeting
	err := handleSOCKSGreeting(conn)
	if err != nil {
		return err
	}
	//handle the SOCKS connection request
	addr,port,err := handleSOCKSConnection(conn)
	if err != nil {
		return err
	}
	//open a connection to the remote host
	connect_addr := net.TCPAddr{IP: addr, Port: port}
	remote_conn,err := net.DialTCP("tcp", nil, &connect_addr)
	if err != nil {
		return err
	}
	//begin communication with client
	client_signal := make(chan int, 0)
	go func(conn net.Conn, remote_conn net.Conn, sig chan int) {
		for {
			buf := make([]byte, 4096)
			n,err := conn.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Println(string(err.Error()))
				break
			}
			_,err = remote_conn.Write(buf[0:n])
			if err != nil {
				fmt.Println(string(err.Error()))
				break
			}
		}
		sig <- 1
	}(conn, remote_conn, client_signal)
	//begin communication with remote host
	remote_signal := make(chan int, 0)
	go func(conn net.Conn, remote_conn net.Conn, sig chan int) {
		for {
			buf := make([]byte, 4096)
			n,err := remote_conn.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Println(string(err.Error()))
				break
			}
			_,err = conn.Write(buf[0:n])
			if err != nil {
				fmt.Println(string(err.Error()))
				break
			}
		}
		sig <- 1
	}(conn, remote_conn, remote_signal)
	<- client_signal //wait for client
	<- remote_signal //wait for remote host
	conn.Close()
	remote_conn.Close()
	return nil
}

/*
* handleSOCKSGreeting (conn net.Conn) error
*
* Helper function to handle the initial greeting part of the SOCKS5 protocol.
* Receives the greeting from the client, validates it, and responds with the server choice.
*/

func handleSOCKSGreeting(conn net.Conn) error {
	//parse greeting
	buf := make([]byte, 2)
	_,err := conn.Read(buf)
	if err != nil {
		return err
	}
	version := int(buf[0])
	if version != SOCKS_VERSION {
		var err errorSOCKS = "Invalid SOCKS version sent by client."
		return err
	}
	nmethods := int(buf[1])
	if nmethods < 1 {
		var err errorSOCKS = "No SOCKS authentication methods supported by client."
		return err
	}
	//read supported methods
	buf = make([]byte, nmethods)
	n,err := conn.Read(buf)
	if n != nmethods {
		var err errorSOCKS = "The nmethods header did not match the number of supported methods sent by the client."
		return err
	}
	//check if unauthenticated connections are supported
	pass := false
	for _,method := range buf {
		if int(method) == 0 {
			pass = true
			break
		}
	}
	if !pass {
		var err errorSOCKS = "The client does not support unauthenticated SOCKS connections."
		return err
	}
	//return server choice
	buf = make([]byte, 2)
	buf[0] = byte(SOCKS_VERSION)
	buf[1] = byte(0)
	_,err = conn.Write(buf)
	return err
}

/*
* handleSOCKSConnection (conn net.Conn) (net.IP, int, error)
*
* Helper function to handle the connection request part of the SOCKS5 protocol.
* Receives the request, validates it and responds.
* This function returns the IP and port that will be required to establish a connection with the remote host.
*/

func handleSOCKSConnection(conn net.Conn) (net.IP, int, error) {
	var addr net.IP
	var port int
	//parse connection request
	buf := make([]byte, 261)
	n,err := conn.Read(buf)
	if err != nil {
		return addr,port,err
	}
	//check version
	version := int(buf[0])
	if version != SOCKS_VERSION {
		var err errorSOCKS = "Invalid SOCKS version sent by client."
		return addr,port,err
	}
	//check cmd
	cmd := int(buf[1])
	if cmd != 1 {
		var err errorSOCKS = "Client requested a command other than CONNECT, which is unsupported."
		return addr,port,err
	}
	//ignore rsv byte, buf[2]
	//check atyp
	atyp := buf[3]
	var index int
	var raw_addr []byte
	switch atyp {
		case 1:
			raw_addr = buf[4:8] //IPv4, 4 bytes
			addr = net.IP(raw_addr)
			index = 8
		case 3:
			addr_size := int(buf[4]) //hostname, 1 to 255 bytes
			raw_addr = buf[5:5+addr_size]
			addr_list,err := net.LookupIP(string(raw_addr))
			if len(addr_list) == 0 {
				var err errorSOCKS = "No IP addresses was found for the hostname requested by client."
				return addr,port,err
			}
			if err != nil {
				return addr,port,err
			}
			addr = addr_list[0]
			index = 5+addr_size
		case 4:
			raw_addr = buf[4:20] //IPv6, 16 bytes
			addr = net.IP(raw_addr)
			index=20
		default:
			var err errorSOCKS = "Invalid address type specified by client."
			return addr,port,err
	}
	//check port
	port_raw := buf[index:index+2]
	port = int(binary.BigEndian.Uint16(port_raw))
	//construct response (we cheat a bit and just modify what we got from the client)
	resp := buf[0:n]
	resp[1] = byte(0)
	_,err = conn.Write(resp)
	return addr,port,err
}

//ERROR DEFINITIONS

type errorSOCKS string

func (err errorSOCKS) Error() string {
	return string(err)
}


