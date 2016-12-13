package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less (b Integer) bool{
	return a <b
}

func (a *Integer) Add(b Integer)  {
	*a += b
}

func main() {
	var a Integer = 1
	if a.Less(2){
		fmt.Println(a,"less 2")
	}
	a.Add(2)
	println(a)
	var arr = [3] int {1,2,3}
	barr:=&arr
	barr[1]++

	fmt.Println(*barr,arr)
	


}