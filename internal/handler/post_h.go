package handler

import (
	"soulvent/internal/dto"
	"soulvent/internal/service"
	"soulvent/internal/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userID := c.GetString("user_id")
	var postCreateReq dto.CreatePostRequest
	if err := c.ShouldBindJSON(&postCreateReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateCreatePost(&postCreateReq, userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreatePost(&postCreateReq, userID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}
	c.JSON(201, postCreateReq)
}

func GetUserPost(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	username := c.Query("username")
	if err != nil {
		c.JSON(400, gin.H{"error": "page and limit should be Integer Value " + err.Error()})
	}

	if err := validators.ValidateGetUserPosts(username, limit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	posts, err := service.GetUserPosts(username, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get posts: " + err.Error()})
	}
	c.JSON(200, posts)
}

func BulkCreatePosts(c *gin.Context) {
	username := c.Query("user_name")
	var postsCreateReq []dto.CreatePostRequest
	if err := c.ShouldBindJSON(&postsCreateReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if err := service.BulkCreatePost(&postsCreateReq, username); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}
	c.JSON(201, postsCreateReq)
}
