package service

import (
	"soulvent/internal/model"
	"soulvent/internal/repository"
)

func CreateUser(user *model.User) error {
	// Logic to create a user in the database
	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}
