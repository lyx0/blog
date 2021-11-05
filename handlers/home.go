package handlers

import "github.com/gin-gonic/gin"

func GetHome(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}
