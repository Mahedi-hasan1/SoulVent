package main

import (
	"log"
	"net/http"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://soulvent-frontend.vercel.app", "https://soulvent.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//unprotected
	r.GET("/", health)
	r.HEAD("/", health)
	r.POST("/users", handler.CreateUser)
	r.DELETE("/feed/clear-old-seen", handler.ClearOldSeenPosts)
	r.POST("/login", handler.Login)
	r.POST("/signup", handler.SignUP)
	r.POST("/posts-bulk", handler.BulkCreatePosts)

	//protected
	protected := r.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		//user routes
		protected.GET("/users", handler.GetUserByUsername)
		protected.GET("/suggested-users", handler.GetSuggestedUsers)
		//post routes
		protected.POST("/posts", handler.CreatePost)
		protected.GET("/posts", handler.GetUserPost)

		//follower routes
		protected.POST("/followers", handler.CreateFollower)
		protected.GET("/followers", handler.GetFollowers)

		//feed routes
		protected.GET("/feed", handler.GetUserFeed)

		//search routes
		protected.POST("/searches", handler.AddSearchHistroy)
		protected.GET("/searches", handler.GetSearchResult)
		protected.GET("search-history", handler.GetSearchHistory)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("SoulVent main service running on :" + port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
