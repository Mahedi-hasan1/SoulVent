package repository

import (
	"soulvent/internal/model"
	"soulvent/internal/db"
)

func CreatePost(post *model.Post) error {
	// Logic to save post to the database
	if err := db.PgDb.Create(post).Error; err != nil {
		return err
	}
	return nil
}