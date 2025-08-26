package main

import (
	"log"
	"soulvent/internal/db"
	"soulvent/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectPostgresDB()
	//db.AutoMigrateModels()
}

func main() {
	r := gin.Default()
	r.POST("/users", handler.CreateUser)
	log.Println("SoulVent main service running on :8080")
	r.Run(":8080")
}
