package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage: %s hostname\n",os.Args[0])
		os.Exit(1)
	}
	_name := os.Args[1]
	_addr,_err := net.LookupHost(_name)

	if _err != nil{
		fmt.Println("Error: ",_err.Error())
		os.Exit(2)
	}

	for _,_s := range _addr{
		fmt.Println(_s)
	}
	os.Exit(0)
}
