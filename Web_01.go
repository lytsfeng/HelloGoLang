package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
)

func sayHelloWeb(response http.ResponseWriter,request *http.Request)  {
	request.ParseForm() // 解析参数
	fmt.Println(request.Form)

	fmt.Println(request.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	response.Header().Set("Content-Length","1024")
	fmt.Fprintf(response, "Hello astaxie!") //这个写入到w的是输出到客户端的
}


func main() {
	http.HandleFunc("/",sayHelloWeb)

	err := http.ListenAndServe(":9999",nil)

	if err != nil{
		log.Fatal("ListenANdServe ",err)
	}



}