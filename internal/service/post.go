package service

import (
	"errors"
	"log"
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/repository"
	"time"
)

func CreatePost(postCreateReq *dto.CreatePostRequest, userID string) error {
	post := &model.Post{
		UserID:    userID,
		Content:   postCreateReq.Content,
		ImageURLs: postCreateReq.ImageURLs,
	}
	return repository.CreatePost(post)
}

func GetUserPosts(userID string, limit int) ([]dto.UserPostResponse, error) {
	return repository.GetUserPosts(userID, limit)
}
func BulkCreatePost(postsCreateReq *[]dto.CreatePostRequest, username string) error {
	now := time.Now()
	users, err := repository.GetUsers("", "", username)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("No User found of this username")
	}

	for i, postCreateReq := range *postsCreateReq {
		createdAt := now.Add(time.Duration(-i*24) * time.Hour)
		post := &model.Post{
			UserID:    users[0].ID,
			Content:   postCreateReq.Content,
			ImageURLs: postCreateReq.ImageURLs,
			CreatedAt: createdAt,
		}
		if err := repository.CreatePost(post); err != nil {
			log.Panicln("post not created. details: ", post)
		} else {
			log.Println("post created: details", post.Content)
		}
	}
	return nil
}
