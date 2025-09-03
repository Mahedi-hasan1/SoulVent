package handler

import (
	"log"
	"net/http"
	"soulvent/internal/service"
	"soulvent/internal/validators"
	"strconv"
	"time"

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


func MarkPostsViewed(c *gin.Context) {
    userID := c.Query("user_id")
    
    var request struct {
        PostIDs []string `json:"post_ids"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    err := service.MarkPostsSeen(userID, request.PostIDs)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark posts as viewed"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func ClearOldSeenPosts(c *gin.Context){
	userId := c.Query("user_id")
	cutoffDate := c.Query("date")
	if cutoffDate == "" {
        cutoffDate = time.Now().Format("2006-01-02 15:04:05")
    }
	
	if err := validators.ValidateClearOldSeenPosts(userId, cutoffDate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := service.ClearOldSeenPost(userId, cutoffDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error()+"Failed to clear old seen posts"})
        return
	}
	 c.JSON(http.StatusOK, gin.H{
        "message": "Old seen posts cleared successfully",
        "cutoff_timestamp": cutoffDate,
        "note": "Posts from the specified date and later were preserved",
    })
}


