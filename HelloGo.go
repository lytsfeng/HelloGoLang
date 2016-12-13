package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func swap(x, y string) (string, string) {
	c, python, java := true, false, "no!"

	fmt.Println(i, j, c, python, java)
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = x - sum
	return
}

var c, python, java bool = true, false, true
var ruby = true

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 2i)
)

func baskType() {

	const f = "%T(%v)\n"
	fmt.Print(f, ToBe)
	fmt.Print(f, MaxInt)
	fmt.Print(f, z)

}

var i, j int = 1, 2

func sqrt(x int32) string {
	if x < 10 {
		return "this is < 10"
	}
	return "this is > 10"
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim{
		return v
	}
	return lim

}

func Compute()(result string, err error){
	result = "hello"

	return;
}

func main() {

	fmt.Print("hello go")
	fmt.Println(rand.Intn(10))
	fmt.Println("%g", math.Nextafter(2, 3))
	fmt.Println(math.Phi)
	fmt.Println(math.Pi)
	fmt.Println(swap("go", "hello"))
	fmt.Println(split(19))
	var i int
	fmt.Println(i, c, python, java, ruby)
	var one, two = true, "hello go"
	three, four := 23, math.E
	fmt.Println(one, two, three, four)
	baskType()
	x, y := 1, 2
	fmt.Println(x, y)

	fmt.Println(sqrt(4), sqrt(14))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	result,err := Compute()
	println(result,err)

}
