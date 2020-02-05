package postex

import "os"
import "net"
import "strconv"
import "io"
import "bytes"

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
	if err != nil {
		return err
	}
	return nil
}

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
	if err != nil {
		return err
	}
	return nil
}
