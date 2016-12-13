/**
dns 解析
Resolution   解析度
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
	_addr,_err := net.ResolveIPAddr("ip",_name)

	if _err != nil{
		fmt.Println("Resolution error",_err.Error())
		os.Exit(1)
	}

	fmt.Println(_name ," Resolution address is ",_addr.String())
	os.Exit(0)
}


