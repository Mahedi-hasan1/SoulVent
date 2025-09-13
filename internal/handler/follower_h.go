package handler

import (
	"fmt"
	"net/http"
	"soulvent/internal/dto"
	"soulvent/internal/service"
	"soulvent/internal/validators"

	"github.com/gin-gonic/gin"
)

func CreateFollower(c *gin.Context) {
	var followerReq dto.CreateFollowerRequest
	if err := c.ShouldBindJSON(&followerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	followerReq.FollowerID = c.GetString("user_id")
	fmt.Println("follower Request :", followerReq)
	if err := validators.ValidateCreateFollower(&followerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateFollower(&followerReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create follower relationship: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, followerReq)
}

func GetFollowers(c *gin.Context) {
	userID := c.GetString("user_id")
	followerID := c.Query("follower_id")
	if err := validators.ValidateGetFollowers(userID, followerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	followers, err := service.GetFollowers(userID, followerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get followers: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, followers)

}
