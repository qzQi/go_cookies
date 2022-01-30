package gee

import (
	"fmt"
	"net/http"
	// "honnef.co/go/tools/pattern"
)

// HandlerFunc_ the proto of http
type HandlerFunc_ func(w http.ResponseWriter, r *http.Request)

// Engine implements the interface of http.handler
type Engine struct {
	router map[string]HandlerFunc_
}

func NewEngine() *Engine {
	return &Engine{make(map[string]HandlerFunc_)}
}

func (e *Engine) addRouter(method, pattern string, handler HandlerFunc_) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) Get(pattern string, handler HandlerFunc_) {
	e.addRouter("Get", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc_) {
	e.addRouter("Post", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 想知道URL里面确切的东西查看一下它的源代码
	key := r.Method + "-" + r.URL.Path
	fmt.Println(key)
	if handler, ok := e.router[key]; ok != false {
		handler(w, r)
	} else {
		fmt.Fprintln(w, "404 NOT FOUND: ", r.URL.Path)
	}
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
