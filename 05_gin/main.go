package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.GET("/api/users", func(c *gin.Context) {
		// Handle GET request for /api/users
	})

	r.POST("/api/users", func(c *gin.Context) {
		// Handle POST request for /api/users
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to my website",
		})
	})

	r.Static("/static", "./static")

	r.Run(":8080")
}
