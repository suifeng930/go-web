package main

import (
	"github.com/go-programming-tour-book/go-web/day05-middleware/hong"
	"log"
	"net/http"
	"time"
)

func main() {

	engine := hong.New()

	engine.Use(hong.Logger())

	engine.GET("/", func(c *hong.Context) {
		c.HTML(http.StatusOK, "<h1>Hello xiao ma </h1>")
	})
	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *hong.Context) {
			c.String(http.StatusOK, "hello, %s you are at %s \n", c.Param("name"), c.Path)
		})
	}

	engine.Run(":9999")
}

func onlyForV2() hong.HandlerFunc {
	return func(c *hong.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}

}
