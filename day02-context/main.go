package main

import (
	"github.com/go-programming-tour-book/go-web/day02-context/hong"
	"net/http"
)

func main() {

	r := hong.New()
	r.GET("/", func(c *hong.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *hong.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *hong.Context) {
		c.JSON(http.StatusOK, hong.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
