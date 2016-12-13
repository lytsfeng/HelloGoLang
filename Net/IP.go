/*
IP 验证
 */
package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage: %s ip-addr\n",os.Args[0])
		os.Exit(1)
	}

	_name := os.Args[1]
	_addr := net.ParseIP(_name)

	if _addr == nil{
		fmt.Println("invalid address")
	}else {
		fmt.Println("the address is ",_addr.String())
	}
	os.Exit(0)
}
