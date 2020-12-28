package main

import (
	"github.com/go-programming-tour-book/go-web/day04-route-group/hong"
	"net/http"
)

func main() {

	engine := hong.New()
	engine.GET("/index", htmlHandler)
	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", stringHandler)
		v1.GET("/hello/:name", helloHandle)

	}
	v2 := engine.Group("v2")
	{
		v2.GET("/assets/*filepath", fileHandler)
		v2.POST("/login", jsonHandler)
	}

	engine.Run(":9999")
}

func htmlHandler(c *hong.Context) {
	c.HTML(http.StatusOK, "<h1>Hi xiao ma</h1>")

}
func stringHandler(c *hong.Context) {

	//
	c.String(http.StatusOK, "hello %s ,you are at %s\n", c.Query("name"), c.Path)

}

func helloHandle(c *hong.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)

}

func fileHandler(c *hong.Context) {
	c.JSON(http.StatusOK, hong.H{"filepath": c.Param("filepath")})
}

func jsonHandler(c *hong.Context) {
	c.JSON(http.StatusOK, hong.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})

}
