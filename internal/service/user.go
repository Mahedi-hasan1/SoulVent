package service

import (
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/repository"
	"soulvent/internal/util"
)

func CreateUser(userReq *dto.CreateUserRequest) error {
	hashedPassword, err := util.HashPassword(userReq.Password)
	if err != nil {
		return err
	}
	user := &model.User{
		Username: userReq.Username,
		Gender: userReq.Gender,
		City:   userReq.City,
		Email:    userReq.Email,
		PasswordHash: hashedPassword,
	}
	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func GetUsers(userId, email, username string) ([]model.User, error){
	return repository.GetUsers(userId, email, username)
}