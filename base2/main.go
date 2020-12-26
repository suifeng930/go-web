package main

import (
	"fmt"
	"log"
	"net/http"
)

// engine is the uni handler for all requests
// 定义 一个空的结构体 Engine
type Engine struct {

}

// implement the nte/http  handler interface
//  实现 ServeHTTP 方法， 两个入参，
// Request : 该对象包含了该HTTP请求的所有的信息（请求地址、Header \Body等信息）
// ResponseWriter  :利用ResponseWriter可以构造针对该请求的响应
func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w,"URL.Path= %q\n",req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w,"Header[%q] =%q\n",k,v)
		}
	default:
		fmt.Fprintf(w,"404 NOT FOUND: %s\n",req.URL)
	}

}
func main() {

	engine := new(Engine)

	log.Fatal(http.ListenAndServe(":9999",engine))

}
