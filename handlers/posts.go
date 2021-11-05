package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type post struct {
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Time  time.Time `json:"time"`
}

var posts = []*post{
	{Title: "First post", Text: "This is the first post", Time: time.Now()},
	{Title: "Second post", Text: "This is the second post", Time: time.Now()},
	{Title: "Third post", Text: "This is the third post", Time: time.Now()},
}

func GetPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}
