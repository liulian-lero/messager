package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoAll)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echoAll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("================")
	fmt.Fprintf(w, "success") // 这个写入到 w 的是输出到客户端的
}