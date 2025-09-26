package db

import (
	"fmt"
	"log"
	"os"
	"soulvent/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgDb *gorm.DB

func ConnectPostgresDB() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("Database URL:", dsn)
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	PgDb = db
}

func AutoMigrateModels() {

	if err := PgDb.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("failed to auto-migrate User Model: ", err)
	} else if err := PgDb.AutoMigrate(&model.Post{}); err != nil {
		log.Fatal("failed to auto-migrate Post Model: ", err)
	} else if err := PgDb.AutoMigrate(&model.Follower{}); err != nil {
		log.Fatal("failed to auto-migrate Follower Model: ", err)
	}else if err := PgDb.AutoMigrate(&model.Search{}); err != nil {
		log.Fatal("failed to auto-migrate Search Model: ", err)
	} else {
		log.Println("Database migrated successfully")
	}
}
