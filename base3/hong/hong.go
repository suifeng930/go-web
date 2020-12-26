package hong

import (
	"fmt"
	"net/http"
)

//  HandlerFunc defines the request  use by hong
// 定义 函数变量
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	// 定义route  用map来映射
	router map[string]HandlerFunc
}

// New is the constructor of hong.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.AddRoute("GET", pattern, handler)
}

// POST	 defines the method to add POST	 request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.AddRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// implements the  net/http Handle
// 解析请求的路径、 查找路由映射表、如果查到、就执行注册的处理方法、如果查不到就返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound) //更新状态码
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
