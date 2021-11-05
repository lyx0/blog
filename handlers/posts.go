package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type post struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	ID    string `json:"id"`
}

var posts = []post{
	{Title: "First Post", Text: "This is the first post", ID: "1"},
	{Title: "Second Post", Text: "This is the second post", ID: "2"},
	{Title: "Third Post", Text: "This is the third post", ID: "3"},
}

func GetPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func AddPosts(c *gin.Context) {
	var p post

	// call BindJSON on the context to parse the request body into the struct
	if err := c.BindJSON(&p); err != nil {
		return
	}
	// Add the new post to the slice
	posts = append(posts, p)
	c.IndentedJSON(http.StatusCreated, posts)
}

func DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")

	var p post
	for i, v := range posts {
		if v.ID == id {
			p = posts[i]
			posts = append(posts[:i], posts[i+1:]...)
			c.IndentedJSON(http.StatusOK, p)
		}
	}
	c.IndentedJSON(http.StatusOK, p)
}
