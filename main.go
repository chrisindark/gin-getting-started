package main

import (
	"net/http"

	"example.com/gin-getting-started/controllers"
	"example.com/gin-getting-started/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
