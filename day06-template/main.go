package main

import (
	"fmt"
	"github.com/go-programming-tour-book/go-web/day06-template/hong"
	"html/template"
	"net/http"
	"time"
)

func main() {
	r :=hong.New()

	r.Use(hong.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate":FormatAdDate,
	})
	r.LoadHTMLGlob("day06-template/templates/*")
	r.Static("/assets","day06-template/static")

	stu1 :=&student{
		Name: "Geektutu",
		Age:  20,
	}
	stu2 :=&student{
		Name: "jack",
		Age:  22,
	}

	r.GET("/", func(c *hong.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *hong.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", hong.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *hong.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", hong.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
	
}

type student struct {
	Name string
	Age int8
}

func FormatAdDate(t time.Time) string {

	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d",year,month,day)

}