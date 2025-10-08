package service

import (
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/repository"
	"soulvent/internal/util"
	"strings"
)

func CreateUser(userReq *dto.CreateUserRequest) error {
	hashedPassword, err := util.HashPassword(userReq.Password)
	if err != nil {
		return err
	}
	username := strings.ReplaceAll(userReq.Username, " ", "_")
	user := &model.User{
		Username:     username,
		Gender:       userReq.Gender,
		City:         userReq.City,
		Email:        userReq.Email,
		PasswordHash: hashedPassword,
	}
	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (dto.UserResponse, error) {
	user, _ := repository.GetUser("", "", username)
	userRes := dto.UserResponse{
		ID:             user.ID,
		Username:       user.Username,
		Gender:         user.Gender,
		Email:          user.Email,
		City:           user.City,
		PostCount:      repository.CountPost(user.ID),
		FollowerCount:  repository.CountFollower(user.ID),
		FollowingCount: repository.CountFollowing(user.ID),
		CreatedAt:      user.CreatedAt,
	}
	return userRes, nil
}

func GetSuggestedUsers(userId string, limit int) ([]model.User, error) {
	return repository.GetSuggestedUsers(userId, limit)
}
