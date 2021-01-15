package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var links = make(map[string]string)
var ids = make(map[string]string)

var port = flag.String("p", "8080", "Port on which to serve")
var file = flag.String("f", "links.txt", "File to which save links")

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	link := links[r.URL.Path[len("/"):]]
	if link == ""{
		home(w, r)
	} else {
		http.Redirect(w, r, link, http.StatusMovedPermanently)
	}
}

func homeOrRedirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		home(w, r)
	} else {
		redirect(w, r)
	}
}

func result(w http.ResponseWriter, check string, correct string, invalid string) {
	t := template.Must(template.ParseFiles("html/result.html"))
	val := ""
	if check == "" {
		val = invalid
	} else {
		val = correct
	}
	err := t.Execute(w, val)
	if err != nil {
		fmt.Println(err)
	}
}

func save(w http.ResponseWriter, r *http.Request) {
	link := r.FormValue("link")
	id := addLink(link)

	result(w, id, r.Host+"/"+id, "invalid id")
}

func decode(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("link")
	link := links[id]

	result(w, link, link, "invalid id")
}

func main() {
	flag.Parse()

	readLinks()

	fmt.Println("ready ", *port)

	http.HandleFunc("/s/", save)
	http.HandleFunc("/d/", decode)
	http.HandleFunc("/", homeOrRedirect)

	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
