package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/lyx0/blog/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.GetHome)
	router.GET("/posts", handlers.GetPosts)
	router.POST("/posts", handlers.AddPosts)
	router.DELETE("/posts", handlers.DeletePost)

	router.Run("localhost:8080")
}
