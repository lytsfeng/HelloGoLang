package main

import "fmt"

type personInfo struct {
	Id      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]personInfo
	personDB = make(map[string]personInfo)

	personDB["1"] = personInfo{"1", "john", "hebei"}
	personDB["2"] = personInfo{"2", "tom", "tj"}

	person, _flag := personDB["3"]

	if _flag {
		println("find", person.Name)
	} else {
		println("no find")
	}
	//v := [] int{1,2}
	//resultChan := make(chan int,2)
	//Sum(v,resultChan)

	_d  := map[int]string{1:"2",2:"4"}
	for k := range _d{
		if k == 1{
			delete(_d,k)
			_d[4] = "hello"
		}
		fmt.Println(k)
	}
	fmt.Println(_d)



	ints := []int{1,2,3,4,5}
	for i := range ints{
		fmt.Println(i)
	}






}
