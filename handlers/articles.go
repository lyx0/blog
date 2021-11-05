package handlers

import (
	"log"
	"net/http"
	"time"
)

type Article struct {
	Heading string    `json:"heading"`
	Body    string    `json:"body"`
	ID      int       `json:"id"`
	Time    time.Time `json:"time"`
	l       *log.Logger
}

func NewArticle(l *log.Logger) *Article {
	return &Article{l: l}
}

func (a *Article) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("ServeHTTP")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("OK"))
}
