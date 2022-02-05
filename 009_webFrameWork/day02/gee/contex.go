package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/qzQi/go_cookies/009_webFrameWork/day01/gee"
)

// H typedef of json
type H map[string]interface{}

// Contex :
type Contex struct {
	ResponseWriter http.ResponseWriter
	Request        http.Request
	Path           string
	Method         string
	StatusCod      int
}

func NewContex(w http.ResponseWriter, r *http.Request) *Contex {
	context := Contex{ResponseWriter: w, Request: *r}
	context.Path = r.URL.Path
	context.Method = r.Method
	// context.
	return &context
}

// PostForm ：只适用于表单数据 k，y。但是post请求可以有消息体啊
// 这个目前没有实现? 看go标准库如何实现的!!!
func (c *Contex) PostForm(key string) string {

	// 就是这里出现了问题，没有解析form
	value := c.Request.FormValue(key)
	if value == "" {
		value = "error in PostForm"
	}
	// fmt.Fprintln(LogFile, "in ", c.Path, "value:", value)
	return value
}

func (c *Contex) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Contex) Status(code int) {
	c.StatusCod = code
	c.ResponseWriter.WriteHeader(code)
}

func (c *Contex) SetHeader(k, v string) {
	c.ResponseWriter.Header().Set(k, v)
}

// String msg-> plain text
func (c *Contex) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.ResponseWriter.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON: msg-> json
// 使用的使用 obj=gee.H --》map[string]interface{}
// for more about json : blog/json for go.com
func (c *Contex) JSON(code int, obj interface{}) {
	// 应该可以这样解析吧，debug时候再看
	b, err := json.Marshal(obj)
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	if err != nil {
		http.Error(c.ResponseWriter, err.Error(), 500)
	}
	// json file本身也是字符串
	c.ResponseWriter.Write([]byte(fmt.Sprintf("%s", b)))
}

func (c *Contex) Data(code int, data []byte) {
	c.Status(code)
	// byte slice 不一定是string，嗯，这样说不对，网络传输都是string 二进制
	c.ResponseWriter.Write(data)
}

func (c *Contex) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	// 前面指出是html
	c.ResponseWriter.Write([]byte(html))
}
