package db

import (
	"log"
	"os"
	"soulvent/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgDb *gorm.DB

func ConnectPostgresDB() {
	dsn := os.Getenv("DATABASE_URL")
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
	err := PgDb.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed to auto-migrate: ", err)
	} else {
		log.Println("Database migrated successfully")
	}
}
