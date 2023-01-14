package main

import (
	"log"
	"net/http"

	gee "github.com/qzQi/go_cookies/009_webFrameWork/day03/gee"
)

func main() {
	e := gee.NewEngine()

	e.GET("/", func(c *gee.Contex) {
		c.String(http.StatusOK, "<h1>Hello world </h1>")
	})

	e.GET("/qzy", func(c *gee.Contex) {
		c.JSON(http.StatusOK, gee.H{
			"user": c.Query("name"),
			"age":  c.Query("age"),
		})
	})

	e.GET("/hello/:name", func(c *gee.Contex) {
		c.String(http.StatusOK, "hello %s, now you are at Path %s\n", c.Param("name"),
			c.Path)
	})

	e.POST("/qzy", func(c *gee.Contex) {
		c.JSON(http.StatusCreated, gee.H{
			"title": c.PostForm("title"),
			"name":  c.PostForm("author"),
		})
	})

	log.Fatal(e.Run("localhost:8080"))

}
