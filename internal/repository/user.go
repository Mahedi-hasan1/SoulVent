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

