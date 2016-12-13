package main

import (
	"fmt"
	"reflect"
)

type Bird struct  {
	Name string
	LifeExpectance int
}

func (b *Bird) Fly(){
	fmt.Println("i am flying "+b.Name)
}

func main() {
	sparrow := &Bird{"Sparrow",3}
	s := reflect.ValueOf(sparrow).Elem()
	typeofT := s.Type()

	for i:=0;i< s.NumField();i++{
		f:=s.Field(i)
		fmt.Printf("%d: %s %s = %v\n",i,typeofT.Field(i).Name,
		f.Type(),f.Interface())
	}
}


