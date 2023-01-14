package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var LogFile = os.Stdout

type H map[string]interface{}

type Contex struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Path           string
	Method         string
	Params         map[string]string
	StatusCode     int
}

func newContex(w http.ResponseWriter, r *http.Request) *Contex {
	return &Contex{
		ResponseWriter: w,
		Request:        r,
		Path:           r.URL.Path,
		Method:         r.Method,
	}
}

func (c *Contex) Param(key string) string {
	v, _ := c.Params[key]
	return v
}

// PostForm: for formData
func (c *Contex) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// Query: for url's query string
func (c *Contex) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// Status:显式调用的话需要在SetHeader之后调用
func (c *Contex) Status(code int) {
	c.StatusCode = code
	c.ResponseWriter.WriteHeader(code)
}

func (c *Contex) SetHeader(k, v string) {
	c.ResponseWriter.Header().Set(k, v)
}

func (c *Contex) String(code int, format string, val ...interface{}) {
	c.SetHeader("Content-Type", "text-plain")
	c.Status(code)
	c.ResponseWriter.Write([]byte(fmt.Sprintf(format, val...)))
}

func (c *Contex) JSON(code int, obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		http.Error(c.ResponseWriter, err.Error(), 500)
		return
	}
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	c.ResponseWriter.Write(b)
}

func (c *Contex) Data(code int, data []byte) {
	c.Status(code)
	c.ResponseWriter.Write(data)
}
