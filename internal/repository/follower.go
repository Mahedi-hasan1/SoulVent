package repository

import (
	"soulvent/internal/db"
	"soulvent/internal/model"
)

func CreateFollower(follower *model.Follower) error {
	if err := db.PgDb.Create(follower).Error; err != nil {
		return err
	}
	return nil
}

func GetFollowers(userID string, followerID string) ([]model.Follower, error) {
	var followers []model.Follower
	query := db.PgDb.Model(&model.Follower{}).Preload("User").Preload("Follower")
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if followerID != "" {
		query = query.Where("follower_id = ?", followerID)
	}
	if err := query.Find(&followers).Error; err != nil {
		return nil, err
	}
	return followers, nil
}