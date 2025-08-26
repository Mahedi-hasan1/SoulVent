package handler

import (
	"net/http"
	"soulvent/internal/model"
	"soulvent/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
	}
	if err := service.CreateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
        return
	}
	 c.JSON(http.StatusCreated, u)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Handle user profile retrieval
}
