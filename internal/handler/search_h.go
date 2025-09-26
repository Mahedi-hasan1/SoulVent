package handler

import (
	"net/http"
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/service"
	"soulvent/internal/validators"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSearchUsers(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID=="" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, err := strconv.Atoi(c.DefaultQuery("page","1"))
	searchReq := &dto.SearchRequest{
		Query: strings.TrimSpace(c.Query("query")),
		Page: page ,
		Limit: limit,
	}
	if err := validators.ValidateSearchRequest(searchReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := service.SearchUsers(userID, searchReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func AddSearchHistroy(c *gin.Context){
	userID := c.GetString("user_id")
	if userID=="" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var addSearchReq  model.Search
	if err := c.ShouldBindJSON(&addSearchReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	addSearchReq.UserID = userID
	if err := validators.ValidateAddSearch(&addSearchReq); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddSearch(&addSearchReq); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, addSearchReq)
}

func GetSearchHistory(c *gin.Context){
	userID := c.GetString("user_id")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	searches, err := service.GetSearches(userID,limit)
	if(err!=nil){
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,searches)
}