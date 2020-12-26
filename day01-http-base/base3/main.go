package main

import (
	"fmt"
	"github.com/go-programming-tour-book/go-web/day01-http-base/base3/hong"
	"net/http"
)

func main() {

	route := hong.New()
	route.GET("/",indexHandler)
	route.GET("/hello",helloHandler)


	route.Run(":9999")
}


//  handler echo os r.url.path
func indexHandler(w http.ResponseWriter,req *http.Request)  {
	fmt.Fprintf(w,"URL.PATH=%q\n",req.URL.Path)
}

//handler echo is req.header 响应的是 请求头header中的键值对信息
func helloHandler(w http.ResponseWriter, req *http.Request)  {
	for k,v :=range req.Header {
		fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
	}

}