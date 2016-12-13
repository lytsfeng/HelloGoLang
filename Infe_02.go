package main

import "fmt"

type Integer int

func (a Integer) MyLess(b Integer) bool  {
	return a < b
}

func (a *Integer) MyAdd(b Integer) {
	*a += b
}

type MyLessAdd interface {
	MyLess(b Integer) bool
	MyAdd(b Integer)
}

type IMyLess interface {
	MyLess(b Integer) bool
}


func main() {
	var a Integer = 1
	var b IMyLess = a
	var c IMyLess = &a

	var d IMyLess = &a 		// 正确
	//var e IMyLess = a    	// 错误


	var iLess MyLessAdd = new(Integer)

	if d ,ok := iLess.(*Integer); ok{
		fmt.Println(d,"-------------------------------")
	}

	var vi interface{} = &a
	switch v := vi.(type) {
	case IMyLess:
		println("---------------&7777777")
	case MyLessAdd:
		println("---------------------88888888888888")
	default:
		println(v)
	}
	println(&a,b,c,d)
}

