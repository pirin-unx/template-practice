package main

import (
	"fmt"
	"net/http"
	"html/template"
)

NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles(basictemplating.html)
	t.execute(w,p)
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is so neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
