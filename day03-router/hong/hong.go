package hong

import (
	"net/http"
)

//  HandlerFunc defines the request  use by hong
// 定义 函数变量
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	// 定义route  用map来映射
	router *router
}

// New is the constructor of hong.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method,pattern,handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST	 defines the method to add POST	 request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// implements the  net/http Handle
// 解析请求的路径、 查找路由映射表、如果查到、就执行注册的处理方法、如果查不到就返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
