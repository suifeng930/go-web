package main

import (
	"github.com/go-programming-tour-book/go-web/day07-panic-recover/hong"
	"net/http"
)

func main() {

	engine := hong.Default()
	engine.GET("/", func(c *hong.Context) {
		c.String(http.StatusOK, "Hello xiao ma \n")
	})
	engine.GET("/panic", func(c *hong.Context) {
		names := []string{"xiaoma"}
		c.String(http.StatusOK, names[100])
	})
	engine.Run(":9999")
}
