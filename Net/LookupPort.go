package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) != 3{
		fmt.Fprintf(os.Stderr,"Usage:%s network-type service \n")
		os.Exit(1)
	}

	_netWorkType := os.Args[1]
	_service := os.Args[2]

	_port,err := net.LookupPort(_netWorkType,_service)
	if err != nil{
		fmt.Println("error : ",err.Error())
		os.Exit(2)
	}

	fmt.Println("service port",_port)
	os.Exit(0)
}
