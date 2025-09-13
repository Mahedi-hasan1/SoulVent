package repository

import (
	"fmt"
	"soulvent/internal/db"
	"soulvent/internal/model"
)



func  CreateUser(user *model.User) error {
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
		fmt.Println("Error in getting users")
		return nil, err
	}
	return users, nil
}

