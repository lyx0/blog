package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/post", post)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

type Post struct {
	Title   string
	Heading string
	Body    string
}

func post(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/post.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	postContent := Post{
		Title:   "My First Post",
		Heading: "My First Post Heading!",
		Body:    `This is my first post, lorem ipsum dolor sit amet`,
	}

	tpl.ExecuteTemplate(rw, "post.gohtml", postContent)

}

func index(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(rw, "index.gohtml", "Home")
}
