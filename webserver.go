package main

import (
	"fmt"
    "net/http"
	"html/template"
	"io/ioutil"
)

var homeTemplate,_ = template.ParseFiles("webview.html")

type StringInfo struct {
	Val string
	Length string
}

type Page struct {
    Title string
	Error string
	Info *StringInfo
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request: %s\n", r.URL.Path);
	fmt.Fprintf(w, "OK")
}


func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request: %s\n", r.URL.Path);

	pg := &Page{Title: "Microservices Test", Info: nil}
	homeTemplate.Execute(w, pg)
}


func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request: %s\n", r.URL.Path);
	s := string(r.URL.Path[len("/info/"):])

	pg := &Page{Title: "String info: " + s, Info: nil}

	url := "http://strlen:8080/v1/len/" + s
	fmt.Printf("svc call: %s\n", url)
	resp, err := http.Get(url)

	if err != nil {
		pg.Error = err.Error()
		homeTemplate.Execute(w, pg)
		return
	}

	if resp.StatusCode != 200 {
		pg.Error = "Unexpected return status from webservice: " + resp.Status
		homeTemplate.Execute(w, pg)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		pg.Error = err.Error()
		homeTemplate.Execute(w, pg)
		return
	}

	pg.Info = &StringInfo{Val: s, Length: string(body)}
	homeTemplate.Execute(w, pg)
}


func main() {
    http.HandleFunc("/", homepageHandler)
    http.HandleFunc("/info/", infoHandler)
    http.HandleFunc("/healthz", healthzHandler)

	fmt.Printf("Starting server on port 8080\n")
    http.ListenAndServe(":8080", nil)
}
