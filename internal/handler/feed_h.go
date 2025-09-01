package handler

import (
	"soulvent/internal/service"
	"soulvent/internal/validators"
	"strconv"
	"log"
	"github.com/gin-gonic/gin"
)

func GetUserFeed(c *gin.Context) {
	userID := c.Query("user_id")
	log.Println("usrId: ", userID)
	pageNum,err:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil{
		c.JSON(400, gin.H{"error": "page and limit should be Integer Value " + err.Error()})
	}
	if err := validators.ValidateGetUserFeed(userID, pageNum, limit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	feed, err := service.GetUserFeed(userID, pageNum, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user feed: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
        "posts": feed,
        "page": pageNum,
        "limit": limit,
        //"has_more": len(feed) == limit,
    })
}
