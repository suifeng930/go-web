package hong

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//给map[string]interface{}起了一个别名hong.H，构建JSON数据时，显得更简洁。

type H map[string]interface{}

// Context目前只包含了http.ResponseWriter和*http.Request，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
type Context struct {

	// origin object
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string
	Params map[string]string //用于解析参数
	// response info
	StatusCode int

	// middleware
	handlers []HandlerFunc
	index    int

	//engine pointer
	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:   req.URL.Path,
		Method: req.Method,
		Req:    req,
		Writer: w,
		index:  -1,
	}

}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)

	}

}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})

}

//提供了访问Query和PostForm参数的方法。
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)

}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)

}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)

}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)

}

// 快速构造String 响应的方法
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))

}

// 快速构造Json 响应的方法
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)

	}

}

//快速构造data的方法
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)

}
// name :模版名称
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	//支持根据模板文件名选择模板进行渲染。
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

//
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value

}
