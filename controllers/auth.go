package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/toby-anderson/cloud-flex/models"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ginc *gin.Context) {
	var input RegisterInput

	if err := ginc.ShouldBindJSON(&input); err != nil {
		ginc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Username = input.Username
	user.Password = input.Password

	_, err := user.Create()

	if err != nil {
		ginc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ginc *gin.Context) {

	var input LoginInput

	if err := ginc.ShouldBindJSON(&input); err != nil {
		ginc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(input.Username, input.Password)

	if err != nil {
		ginc.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	ginc.JSON(http.StatusOK, gin.H{"token": token})

}
