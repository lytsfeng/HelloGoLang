package main

import (
	"fmt"
)
func main() {
	fmt.Println(65)

	str := "Hello golang"
	_a := [10]int{10,10}
	_b := []int{10,10,01,}
	fmt.Println(_a,len(&_a),cap(_a),&_b)
	for _k,_v:= range str{
		fmt.Println(_k,string(rune(_v)))
	}

	fmt.Println(string(rune('中')))

	for _ss,_v := range []rune("中国golang"){
		fmt.Println(_ss,string(_v))
	}

	fmt.Println(len("go编程"))
	fmt.Println(len("编"))
	fmt.Println(len([]rune("go编程")))

	for _k,_v := range []rune("go编程"){
		fmt.Println(_k,string(_v))
		//switch _v.(type) {
		//case string:
		//
		//}
	}


	Module2(Module(2,4))


}

func Module(x,y int ) (result int,ok bool)  {
	result = x % y
	return
}
func Module2 (x int,ok bool)  {
	fmt.Println(x,"    ",ok)
}


