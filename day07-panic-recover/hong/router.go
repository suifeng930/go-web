package hong

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node       // 存储每种请求方式的Trie 树根节点。
	handlers map[string]HandlerFunc //存储每种请求方式的handlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析路由规则，返回一个路由数组
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {

	// 将路径注册到一个数组中
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok { // 如果没有找到这个请求方法，为其注册一个 root 根树
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0) //注册到根节点中
	r.handlers[key] = handler

}

//解析了:和*两种匹配符的参数，返回一个 map 。
//例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}，
///static/css/geektutu.css匹配到/static/*filepath，解析结果为{filepath: "css/geektutu.css"}。
func (r *router) getRoute(method, path string) (*node, map[string]string) {

	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method] //查找根树
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes

}
func (r *router) handle(c *Context) {

	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		//将从路由个匹配得到的handler 添加到 c.handlers 列表中
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(context *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next() //调用c.next()
}
