package main

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get()  {

	this.Ctx.WriteString("Hello world")
}

func init() {
	fmt.Println("main init")
}

type Test struct {
	name string
}

func (T Test) init()  {
	fmt.Println(T.name)
}

func (T Test) toString()  {
	fmt.Println("this is toString ",T.name)
}





func main() {

	_t := Test{"john"}
	_t.toString()

	beego.Router("/",&MainController{})
	beego.Run()
}
