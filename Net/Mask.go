/**
mask 掩码操作
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
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	_mask := _addr.DefaultMask()
	_newWork :=_addr.Mask(_mask)
	_ones,_bits := _mask.Size()

	fmt.Println("Address is ",_addr.String(),
		"\nDefault mask length is ",_bits,
		"\nLeading ones count is ",_ones,
		"\nMask is (hex) ",_mask.String(),
		"\nnetwork is ",_newWork)
	os.Exit(0)
}


