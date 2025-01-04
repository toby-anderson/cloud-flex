package main

import (
	"github.com/gin-gonic/gin"
	"github.com/toby-anderson/cloud-flex/controllers"
	"github.com/toby-anderson/cloud-flex/models"
)

func main() {
	models.ConnectDataBase()
	ginServe := gin.Default()
	ginServe.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "The web server is running"})
	})
	ginServe.POST("/register", controllers.Register)
	ginServe.Run(":8080")
}
