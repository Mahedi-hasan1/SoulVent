package handler

import (
	"net/http"
	"soulvent/internal/dto"
	"soulvent/internal/service"
	"soulvent/internal/validators"

	"github.com/gin-gonic/gin"
)

func SignUP(c *gin.Context) {
	var signUpReq dto.CreateUserRequest
	if err := c.ShouldBindJSON(&signUpReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateCreateUser(&signUpReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, token, err := service.SignUP(signUpReq.Username, signUpReq.Email, signUpReq.Password, signUpReq.Gender, signUpReq.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func Login(c *gin.Context) {
	var logInReq dto.LogInReq
	if err := c.ShouldBindJSON(&logInReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidateLogIn(&logInReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := service.LogIn(logInReq.UsernameOrEmail, logInReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
