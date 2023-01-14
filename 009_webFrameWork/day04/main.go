package main

import (
	"net/http"

	gee "github.com/qzQi/go_cookies/009_webFrameWork/day04/gee"
)

func main() {
	e := gee.NewEngine()

	e.GET("/", func(c *gee.Contex) {
		c.Data(http.StatusOK, []byte("hello in /index the same / ?. path:"+c.Path))
	})

	g1 := e.NewGroup("/qzy")
	// /qzy or /qzy/
	g1.GET("/", func(c *gee.Contex) {
		c.String(http.StatusOK, "in path:%s\n name:%s\n", c.Path, c.Query("name"))
	})

	// /qzy/lover  no handler
	g2 := g1.NewGroup("/lover")

	//  /qzy/lover/zbx?
	g2.GET("/zbx", func(c *gee.Contex) {
		c.String(http.StatusOK, "name=%s\n age=%s\n", c.Query("name"), c.Query("age"))
	})

	e.Run("localhost:8080")

}
