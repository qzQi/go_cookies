package main

import (
	"fmt"
	"net/http"

	gee "github.com/qzQi/go_cookies/009_webFrameWork/day01/gee"
)

func main() {
	r := gee.NewEngine()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "the r.URL is %#v", r.URL)
	})

	r.Get("/qzy", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] =%q] \n", k, v)
		}
	})

	r.Run("localhost:8080")
}
