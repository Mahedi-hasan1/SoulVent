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

func GetPosts(postID string, userID string) ([]model.Post, error) {
	var posts []model.Post
	query := db.PgDb.Model(&model.Post{})
	if postID != "" {
		query = query.Where("id = ?", postID)
	}
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}