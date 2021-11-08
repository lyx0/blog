package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/post", post)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":8080", nil)

}

func post(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/post.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(rw, "post.gohtml", " Good Post title")

}
