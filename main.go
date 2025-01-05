package main

import (
	"github.com/gin-gonic/gin"
	"github.com/toby-anderson/cloud-flex/controllers"
	"github.com/toby-anderson/cloud-flex/middleware"
	"github.com/toby-anderson/cloud-flex/models"
)

func main() {
	models.ConnectDataBase()
	ginServe := gin.Default()
	publicServe := ginServe.Group("/v1")
	protectedServe := ginServe.Group("/v1")
	protectedServe.Use(middleware.JwtAuthHandler())

	publicServe.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "The web server is running"})
	})
	
	publicServe.POST("/register", controllers.Register)
	publicServe.POST("/login", controllers.Login)

	protectedServe.GET("/account", controllers.CurrentUser)

	ginServe.Run(":8080")
}
