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

func GetPost(c *gin.Context) {
	postID := c.Query("id")
	userID := c.Query("user_id")
	if err := validators.ValidateGetPosts(postID, userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	posts, err := service.GetPosts(postID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get posts: " + err.Error()})
	}
	c.JSON(200, posts)
}

func BulkCreatePosts(c *gin.Context){
	var postsCreateReq  []dto.CreatePostRequest
	if err := c.ShouldBindJSON(&postsCreateReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if err := service.BulkCreatePost(&postsCreateReq); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}
	c.JSON(201, postsCreateReq)
}