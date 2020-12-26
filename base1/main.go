package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)


	// handler = nil 表示：使用标准库中对实例处理。第二个参数，则是我们给予net/http标准库实现web框架对入口。
	log.Fatal(http.ListenAndServe(":9999",nil))

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