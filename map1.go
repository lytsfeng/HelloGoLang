package main

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
}
