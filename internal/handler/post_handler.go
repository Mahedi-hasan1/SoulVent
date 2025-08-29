package handler

import (
	"soulvent/internal/dto"
	"soulvent/internal/service"
	"soulvent/internal/validators"
	"github.com/gin-gonic/gin"
	"fmt"
)

func CreatePost(c *gin.Context) {
	fmt.Println("new create post request")
	var postCreateReq  dto.CreatePostRequest
	if err := c.ShouldBindJSON(&postCreateReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateCreatePost(&postCreateReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreatePost(&postCreateReq); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}
	c.JSON(201, postCreateReq)
}