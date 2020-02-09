package postex
//Contains functions for sending and receiving files.

import "os"
import "net"
import "crypto/tls"
import "strconv"
import "io"
import "bytes"

/*
* SendFile(filepath string, host string, port int) error
* Open a TCP connection to a host and port; send the specified file along it.
*/

func SendFile(filepath string, host string, port int) error {
	//check if the file exists
	info,err := os.Stat(filepath)
	if err != nil {
		return err
	}
	//attempt to open the file
	fd,err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()
	//read from the file
	filesize := info.Size()
	buf := make([]byte, filesize)
	_,err = fd.Read(buf)
	if err != nil {
		return err
	}
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	target_addr,err := net.ResolveTCPAddr("tcp4", addr_string)
	if err != nil {
		return err
	}
	target_conn,err := net.DialTCP("tcp4", nil, target_addr)
	if err != nil {
		return err
	}
	defer target_conn.Close()
	//send the file
	_,err = target_conn.Write(buf)
	return err
}

/*
* RecvFile(filepath string, host string, port int) error
* Open a TCP connection to a host and port; receive a file on it and write it to the specified path.
*/

func RecvFile(filepath string, host string, port int) error {
	//attempt to open the file
	fd,err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer fd.Close()
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	target_addr,err := net.ResolveTCPAddr("tcp4", addr_string)
	if err != nil {
		return err
	}
	target_conn,err := net.DialTCP("tcp4", nil, target_addr)
	if err != nil {
		return err
	}
	defer target_conn.Close()
	//read in the file and write it locally
	out := bytes.Buffer{}
	reading := true
	for reading {
		buf := make([]byte, 500)
		read_count,err := target_conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				reading = false
			} else {
				return err
			}
		}
		fragment := make([]byte, read_count)
		copy(fragment, buf[:read_count])
		out.Write(fragment)
	}
	out_bytes := out.Bytes()
	_,err = fd.Write(out_bytes)
	return err
}

/*
* SendFileTLS(filepath string, host string, port int, skip_verify bool) error
* Open a TCP connection to a host and port, encrypted with TLS; send the specified file along it.
*/

func SendFileTLS(filepath string, host string, port int, skip_verify bool) error {
	//check if the file exists
	info,err := os.Stat(filepath)
	if err != nil {
		return err
	}
	//attempt to open the file
	fd,err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()
	//read from the file
	filesize := info.Size()
	buf := make([]byte, filesize)
	_,err = fd.Read(buf)
	if err != nil {
		return err
	}
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	conf := &tls.Config{InsecureSkipVerify: skip_verify}
	target_conn,err := tls.Dial("tcp4", addr_string, conf)
	if err != nil {
		return err
	}
	defer target_conn.Close()
	//send the file
	_,err = target_conn.Write(buf)
	return err
}

/*
* RecvFileTLS(filepath string, host string, port int, skip_verify bool) error
* Open a TCP connection to a host and port, encrypted with TLS; receive a file on it and write it to the specified path.
*/

func RecvFileTLS(filepath string, host string, port int, skip_verify bool) error {
	//attempt to open the file
	fd,err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer fd.Close()
	//set up the socket
	addr_string := host + ":" + strconv.Itoa(port)
	conf := &tls.Config{InsecureSkipVerify: skip_verify}
	target_conn,err := tls.Dial("tcp4", addr_string, conf)
	if err != nil {
		return err
	}
	defer target_conn.Close()
	//read in the file and write it locally
	out := bytes.Buffer{}
	reading := true
	for reading {
		buf := make([]byte, 500)
		read_count,err := target_conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				reading = false
			} else {
				return err
			}
		}
		fragment := make([]byte, read_count)
		copy(fragment, buf[:read_count])
		out.Write(fragment)
	}
	out_bytes := out.Bytes()
	_,err = fd.Write(out_bytes)
	return err
}

