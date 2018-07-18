package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	r.ParseForm() //解析参数，默认是不会解析的
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}
func sayHello2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	r.ParseForm() //解析参数，默认是不会解析的
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}
func main() {
	http.HandleFunc("/hello", sayHello)   //注册URI路径与相应的处理函数
	http.HandleFunc("/hello2", sayHello2)
	er := http.ListenAndServe(":9090", nil) // 监听9090端口，就跟javaweb中tomcat用的8080差不多一个意思吧
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}
}
