package main

import (

	"fmt"
	"strconv"
	"os"
	"bytes"
	"encoding/binary"
)

type Freq struct {
	With int32
//	Date []byte
}


func main() {
	fmt.Println(strconv.ParseInt("29360640",16,0))
	num2str := 223
	fmt.Println(num2str,"")

	fmt.Fprintf(os.Stdout,"%X \n","h e l l o golang")


	fmt.Println(string([]byte("hello golang")))
	fmt.Println([]rune("hello golang"))
	fmt.Println(string([]byte{104, 101, 108, 108, 111, 32, 103, 111, 108, 97, 110, 103}))
	fmt.Println(string([]rune{20014,22269}))

	for i:=20000;i < 20010;i++{
		fmt.Println(string((i)))
	}

	fmt.Println([]rune("张亚峰"))


	defer func() {
		recover()
		fmt.Println("------------------------")
	}()

	_freq :=  Freq{34}
	var _buf = &bytes.Buffer{}

	err := binary.Write(_buf,binary.BigEndian,_freq)
	if err != nil{
		fmt.Println(err.Error())
		//panic(err)
	}

	fmt.Printf("%d\n",_buf.Bytes())
	fmt.Println(_buf.Bytes())
//	fmt.Printf("%x",sha1.Sum(_buf.Bytes()))
	fmt.Println(cap(_buf.Bytes()))
	_freq2 := Freq{}
	binary.Read(_buf,binary.BigEndian,&_freq2)
	fmt.Println(_freq2.With)

}

