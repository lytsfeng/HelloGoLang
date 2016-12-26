package main

import (
	"HelloGoLang/utils"
	"fmt"
)

func main() {
	var set utils.Set = utils.NewHashSet()
	set.Add("hhh")
	set.Add(true)

	fmt.Println(set.Contains(true))
	fmt.Println(set.Elements())

}
