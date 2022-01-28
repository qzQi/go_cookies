/*
Creating a data structure with load and save methods
Using the net/http package to build web applications
Using the html/template package to process HTML templates
Using the regexp package to validate user input
Using closures
https://go.dev/doc/articles/wiki/
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{fileName, body}, nil
}

// http的querystring，以前不知道
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "in / func")
	fmt.Fprintln(w, r.URL.Path)
	fmt.Fprintf(w, "hi there,i love %s.\n", r.URL.Path[1:])
}

// localhost:8080/view/qzy
func viewHandler(w http.ResponseWriter, r *http.Request) {

	// title := r.URL.Path[len("/view/"):]
	// p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// fmt.Fprintln(w, "in /view/ router")
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		// 原文件不存在的话，就转向edit创建
		// 重定向
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
	}
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

// handler for edit
func editHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/edit/"):]
	// p, err := loadPage(title)
	// if err != nil {
	// 	p = &Page{Title: title}
	// }
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)
	// // This function will work fine,
	// // but all that hard-coded HTML is ugly.
	// // Of course, there is a better way.
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

// after edit we need save the file
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// my test query
func queryStr(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in query", " ", r.URL.RawQuery)
	fmt.Fprintln(w, r.URL.RawQuery)
}

func main() {
	// http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/query", queryStr)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
