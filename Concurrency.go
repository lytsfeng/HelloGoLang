package main

import "fmt"
// 并发


var c chan string

func pingpang()  {
	i := 0
	for{
		fmt.Println(<- c)
		c <- fmt.Sprintf("from pingpang, Hi #%d ",i)
		i++
	}
}

func fibonaccin(n int, c chan int)  {
	x,y :=1,1
	for i :=0;i < n;i++{
		 c <- x
		 x,y = y ,x+y
	}
	close(c)
}


func main() {
	c = make(chan  string)
	go  pingpang();
	for i := 0; i < 10 ; i++  {
		c <- fmt.Sprintf("from main, Hi #%d ",i)
		fmt.Println(<- c)
	}

	c := make(chan int,10 )

	go fibonaccin(cap(c),c)


	for i:= range c{
		fmt.Println(i)
	}





}

