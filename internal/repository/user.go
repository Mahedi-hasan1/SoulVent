package repository

import (
	"soulvent/internal/db"
	"soulvent/internal/model"
)

func CreateUser(user *model.User) error {
	if err := db.PgDb.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers(userId, email, username string) ([]model.User, error) {
	var users []model.User
	query := db.PgDb.Model(&model.User{})

	if userId != "" {
		query = query.Where("id = ?", userId)
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if username != "" {
		query = query.Where("username = ?", username)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetSuggestedUsers(userId string, limit int) ([]model.User, error) {

	var suggestedUsers []model.User
	err := db.PgDb.
		Table("posts").
		Select("users.*").
		Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.user_id != ?", userId).
		Where("posts.user_id NOT IN (?)",
			db.PgDb.Table("followers").
				Select("user_id").
				Where("follower_id = ?", userId),
		).
		Group("users.id").
		Order("COUNT(posts.id) DESC").
		Limit(limit).
		Find(&suggestedUsers).Error

	if err != nil {
		return nil, err
	}

	return suggestedUsers, nil
}
