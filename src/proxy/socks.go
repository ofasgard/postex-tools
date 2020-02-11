package main
//package proxy
//Contains functions for implementing a SOCKS proxy server in Go.

import "net"
import "strconv"
import "fmt"
import "encoding/binary"

//IT WORKS!
//but it needs a lot more testing

const SOCKS_VERSION = 5

func StartProxy(port int) {
	port_str := ":" + strconv.Itoa(port)
	server_addr,err := net.ResolveTCPAddr("tcp4", port_str)
	if err != nil {
		fmt.Println(string(err.Error()))
		return
	}
	listener,err := net.ListenTCP("tcp4", server_addr)
	if err != nil {
		fmt.Println(string(err.Error()))
		return
	}
	for {
		conn,err := listener.Accept()
		if err == nil {
			err := handleSOCKS(conn)
			if err != nil {
				fmt.Println(string(err.Error()))
			}
			conn.Close()
		} else {
			fmt.Println(string(err.Error()))
		}
	}
}

//FUNCTIONS FOR HANDLING SOCKS

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
	defer remote_conn.Close()
	//begin communication with client
	go func(conn net.Conn, remote_conn net.Conn) {
		for {
			buf := make([]byte, 4096)
			n,err := conn.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Println(string(err.Error()))
			}
			_,err = remote_conn.Write(buf[0:n])
			if err != nil {
				fmt.Println(string(err.Error()))
			}
		}
	}(conn, remote_conn)
	//begin communication with remote host
	for {
		buf := make([]byte, 4096)
		n,err := remote_conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(string(err.Error()))
		}
		_,err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Println(string(err.Error()))
		}
	}
	return nil
}

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
	//construct response
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

//TEMPORARY MAIN FUNCTION

func main() {
	//this is temporary
	StartProxy(1080)
}

//https://rushter.com/blog/python-socks-server/
//curl --socks5-hostname localhost:1080 http://www.google.com/
//https://en.wikipedia.org/wiki/SOCKS
