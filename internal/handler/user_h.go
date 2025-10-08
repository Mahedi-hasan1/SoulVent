package handler

import (
	"fmt"
	"net/http"
	"soulvent/internal/dto"
	"soulvent/internal/service"
	"soulvent/internal/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userReq dto.CreateUserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateCreateUser(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateUser(&userReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userReq)
}

func GetUserByUsername(c *gin.Context) {
	username := c.Query("username")
	fmt.Println("username; ", username)
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username required"})
		return
	}
	users, err := service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetSuggestedUsers(c *gin.Context) {
	userId := c.GetString("user_id")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "3"))
	suggestedUsers, err := service.GetSuggestedUsers(userId, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, suggestedUsers)
}
