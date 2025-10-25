package handler

import (
	"net/http"
	"soulvent/internal/model"
	"soulvent/internal/service"
	"soulvent/internal/validators"

	"github.com/gin-gonic/gin"
)

func AddReaction(c *gin.Context) {
	userID := c.GetString("user_id")
	var reactionAddReq *model.Reaction
	if err := c.ShouldBindJSON(&reactionAddReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request" + err.Error()})
		return
	}
	reactionAddReq.UserID = userID
	if err := validators.ValidateAddReaction(reactionAddReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddReaction(reactionAddReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reactionAddReq)
}

func RemoveReaction(c *gin.Context) {
	userID := c.GetString("user_id")
	postID := c.Query("post_id")
	if userID == "" || postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and post_id is required"})
		return
	}
	if err := service.RemoveReaction(postID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"post_id": postID, "user_id": userID})
}
