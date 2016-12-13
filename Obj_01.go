package main

import (
	"fmt"
	"log"
)

type Rect struct {
	x,y float64
	width,height float64
	log.Logger
}

func (r *Rect) area() float64  {
	return r.height * r.width
}

type job struct {
	*Rect
}


func main() {
	r := &Rect{width:200,height:400}
	area := r.area()
	fmt.Println(area)
	fmt.Println(r)
	fmt.Println(*r)
	job{}.area()
}
