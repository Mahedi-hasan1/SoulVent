package main

import (
	"log"
	"soulvent/internal/db"
	"soulvent/internal/handler"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func init() {
	db.ConnectPostgresDB()
	//db.AutoMigrateModels()
	db.InitRedis()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	//user routes
	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.GetUser)

	//post routes
	r.POST("/posts", handler.CreatePost)
	r.POST("/posts-bulk", handler.BulkCreatePosts)
	r.GET("/posts", handler.GetPost)

	//follower routes
	r.POST("/followers", handler.CreateFollower)
	r.GET("/followers", handler.GetFollowers)

	//feed routes
	r.GET("/feed", handler.GetUserFeed)
	r.DELETE("/feed/clear-old-seen",handler.ClearOldSeenPosts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("SoulVent main service running on :" + port)
	r.Run(":" + port)
}
