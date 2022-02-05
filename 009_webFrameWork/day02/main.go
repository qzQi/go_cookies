package main

import (
	"net/http"

	gee "github.com/qzQi/go_cookies/009_webFrameWork/day02/gee"
)

func main() {
	r := gee.NewEngine()
	// usage: /qzy?name=qizhiyun&age=21
	r.GET("/qzy", func(c *gee.Contex) {
		// c.String(http.StatusOK, "hello %s, you are %s year old\n",
		// 	c.Query("name"), c.Query("age"))

		// 这次发送get请求把，内容放在body里面，测试一下json
		c.JSON(http.StatusOK, gee.H{
			"name": c.PostForm("name"),
			"age":  c.PostForm("age"),
		})

		// 测试一下是不是json的解析,就是解析postform 出现了错误
		// json的marshal没问题
		// c.JSON(http.StatusOK, gee.H{
		// 	"name": c.Query("name"),
		// 	"age":  c.Query("age"),
		// })
	})

	// usage: post /login
	// 目前post不可以用
	r.POST("/login", func(c *gee.Contex) {
		c.JSON(http.StatusOK, gee.H{
			"user": c.PostForm("username"),
			"psw":  c.PostForm("password"),
		})
	})

	r.Run("localhost:8080")
}
