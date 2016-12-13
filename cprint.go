package main

/*
#include<stdio.h>
*/

import (
	"C"
	"fmt"
)





func main()  {
	//cstr := C.CString("Hello,World")
	//C.puts(cstr)
	//C.free(unsafe.Pointer(cstr))


	f := func(a,b int) int{
		return a+b
	}

	fmt.Println(f(2,3))

}

