package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	service := ":1201"
	tcpaddr,err := net.ResolveTCPAddr("tcp4",service)
	if err != nil{
		fmt.Fprintf(os.Stderr,"fatal error : ",err.Error())
		os.Exit(1)
	}
	listener,err := net.ListenTCP("tcp4",tcpaddr)
	if err != nil{
		fmt.Fprintf(os.Stderr,"fatal error : ",err.Error())
		os.Exit(1)
	}
	for{
		conn,err := listener.Accept()
		if err != nil{
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	for{
		var buf[512]byte

		n,err:=  conn.Read(buf[0:])

		if err != nil{
			continue
		}

		result := string(buf[0:n-2])
		fmt.Println(result)

		if result == "quit" || result == "exit"{
			break
		}
		_,err = conn.Write(buf[0:n])
		if err != nil{
			continue
		}
	}
}



