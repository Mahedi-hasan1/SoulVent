package repository

import (
	"soulvent/internal/model"
	"soulvent/internal/db"
)



func  CreateUser(user *model.User) error {
	if err := db.PgDb.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers(userId, email string) ([]model.User, error) {
	var users []model.User
	query := db.PgDb.Model(&model.User{})

	if userId != "" {
		query = query.Where("id = ?", userId)
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

