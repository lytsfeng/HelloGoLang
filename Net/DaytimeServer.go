package main

import (
	"net"
	"time"
	"os"
	"fmt"
)

func main() {
	_service := "127.0.0.1:1200"
	_tcpAddr,err := net.ResolveTCPAddr("tcp4",_service)
	if err != nil{
		fmt.Fprintf(os.Stderr,"Fatal error:%s ",err.Error())
		os.Exit(1)
	}
	_listener,err := net.ListenTCP("tcp",_tcpAddr)

	for{
		_conn, err:= _listener.Accept()
		if err != nil{
			continue
		}
		_daytime := time.Now().String()
		_conn.Write([]byte(_daytime))
		_conn.Close()
	}
}
