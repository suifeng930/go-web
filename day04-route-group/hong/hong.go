package hong

import (
	"log"
	"net/http"
)

//  HandlerFunc defines the request  use by hong
// 定义 函数变量
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc //support middleware
		parent      *RouterGroup  //support nesting
		engine      *Engine       //all groups share a Engine instance
	}
	Engine struct { //将Engine作为最顶层的分组，也就是说Engine拥有RouterGroup所有的能力
		*RouterGroup
		// 定义route  用map来映射
		router *router
		groups []*RouterGroup // store all groups
	}
)

// New is the constructor of hong.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}             //初始化构造router 对象
	engine.RouterGroup = &RouterGroup{engine: engine}  //初始化构造RouterGroup对象
	engine.groups = []*RouterGroup{engine.RouterGroup} //将routerGroup 追加到 分组对象中
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {

	engine := group.engine    //获取到engine对象
	newGroup := &RouterGroup{ //初始化一个routerGroup对象
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	//将newGroup对象注入到 routerGroup数组中
	engine.groups = append(engine.groups, newGroup)
	return newGroup

}

//调用了group.engine.router.addRoute来实现了路由的映射。
//由于Engine从某种意义上继承了RouterGroup的所有属性和方法，因为 (*Engine).engine 是指向自己的。
//这样实现，我们既可以像原来一样添加路由，也可以通过分组添加路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {

	pattern := group.prefix + comp
	log.Printf("Route %4s - %s \n", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST	 defines the method to add POST	 request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
