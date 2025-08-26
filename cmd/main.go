package main

import (
	"log"
	"net/http"
	"os"
	"soulvent/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Use DATABASE_URL from .env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Auto-migrate User model
	err = db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal("failed to auto-migrate: ", err)
	} else {
		log.Println("Database connected and migrated successfully")
	}
	// Set up repository for handler

	//http.HandleFunc("/users", handler.CreateUser)

	log.Println("SoulVent main service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
