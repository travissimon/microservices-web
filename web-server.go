package main

import (
    "net/http"
	"html/template"
)

var homeTemplate, _ = template.ParseFiles("web-view.html")

type StringInfo struct {
	Val string
	Length string
}

type Page struct {
    Title string
	Info *StringInfo
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	pg := &Page{Title: "Microservices Test", Info: nil}
	homeTemplate.Execute(w, pg)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	pg := &Page{Title: "String info:", Info: nil}
	homeTemplate.Execute(w, pg)
}


func main() {
    http.HandleFunc("/", homepageHandler)
    http.HandleFunc("/info/", infoHandler)
    http.ListenAndServe(":8080", nil)
}
