package main

import (
    "net/http"
	"html/template"
)

var homeTemplate, _ = template.ParseFiles("web-view.html")

type Page struct {
    Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	pg := &Page{Title: "Microservices Test"}
	homeTemplate.Execute(w, pg)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
