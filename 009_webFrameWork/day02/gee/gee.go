package gee

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Engine struct {
	router *router
}

var LogFile = os.Stdout

func NewEngine() *Engine {
	return &Engine{newRouter()}
}

func (e *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	fmt.Fprintf(LogFile, "Router %4s - %s \n", method, pattern)
	e.router.addRouter(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

// ServeHTTP: control the default ServeMutex(nil in ListenAndServe)
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContex(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}
