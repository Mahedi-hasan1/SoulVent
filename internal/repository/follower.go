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

func GetFollowingIDs(userID string) ([]string, error) {
	var followingIDs []string

	query := db.PgDb.Model(&model.Follower{}).Select("user_id")
	if userID != "" {
		query = query.Where("follower_id = ?", userID)
	}

	if err := query.Pluck("user_id", &followingIDs).Error; err != nil {
		return nil, err
	}

	return followingIDs, nil
}

func IsFollowing(currentUserID, targetUserID string) bool {
	if currentUserID == targetUserID {
		return false
	}
	var count int64
	db.PgDb.Model(&model.Follower{}).
		Where("user_id = ? AND follower_id = ?", targetUserID, currentUserID).
		Count(&count)
	return count > 0
}

func CountFollower(userID string) int64 {
	var count int64
	db.PgDb.Model(&model.Follower{}).Where("user_id = ?", userID).Count(&count)
	return count
}

func CountFollowing(userID string) int64 {
	var count int64
	db.PgDb.Model(&model.Follower{}).Where("follower_id = ?",userID).Count(&count)
	return count
}