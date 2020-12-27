package main

import (
	"github.com/go-programming-tour-book/go-web/day03-router/hong"
	"net/http"
)

func main() {

	engine := hong.New()
	engine.GET("/", htmlHandler)
	engine.GET("/hello", stringHandler)
	engine.GET("/hello/:name", helloHandle)
	engine.GET("/assets/*filepath", fileHandler)

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
