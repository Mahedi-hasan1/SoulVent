package main

import (
	"log"
	"os"
	"soulvent/internal/db"
	"soulvent/internal/handler"
	"soulvent/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectPostgresDB()
	//db.AutoMigrateModels()
	db.InitRedis()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	//unprotected
	r.POST("/users", handler.CreateUser)
	r.DELETE("/feed/clear-old-seen", handler.ClearOldSeenPosts)
	r.POST("/login", handler.Login)
	r.POST("/signup", handler.SignUP)

	protected := r.Group("")

	protected.Use(middleware.AuthMiddleware())
	{
		//user routes
		protected.GET("/users", handler.GetUser)
		//post routes
		protected.POST("/posts", handler.CreatePost)
		protected.POST("/posts-bulk", handler.BulkCreatePosts)
		protected.GET("/posts", handler.GetPost)

		//follower routes
		protected.POST("/followers", handler.CreateFollower)
		protected.GET("/followers", handler.GetFollowers)

		//feed routes
		protected.GET("/feed", handler.GetUserFeed)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("SoulVent main service running on :" + port)
	r.Run(":" + port)
}
