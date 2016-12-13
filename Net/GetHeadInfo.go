package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

func checkError(err error){
	if err != nil{
		fmt.Fprint(os.Stderr,"Fatal error : %s", err.Error())
		os.Exit(1)
	}
}


func main() {
	if len(os.Args) != 2{
		fmt.Fprint(os.Stderr,"Usage: %s host:port ",os.Args[0])
		os.Exit(1)
	}

	_ServiceName := os.Args[1]

	//通过名称拿到地址

	_addr,err := net.ResolveTCPAddr("tcp4",_ServiceName)
	checkError(err)

	_conn,err := net.DialTCP("tcp",nil,_addr)
	checkError(err)

	_ ,err = _conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	_result,err := readFully(_conn)
	checkError(err)
	fmt.Println(string(_result))
	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte,error) {
	defer func() {
		conn.Close()
	}()

	reuslt := bytes.NewBuffer(nil)
	var buf [512]byte

	for{
		n,err := conn.Read(buf[0:])
		reuslt.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return reuslt.Bytes(),nil





}


