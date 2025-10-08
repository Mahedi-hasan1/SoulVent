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

func GetUserPosts(username string, limit int) ([]dto.UserPostResponse, error) {
	return repository.GetPostsByUsername(username, limit)
}
func BulkCreatePost(postsCreateReq *[]dto.CreatePostRequest, username string) error {
	now := time.Now()
	user, err := repository.GetUser("", "", username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("No User found of this username")
	}

	for i, postCreateReq := range *postsCreateReq {
		createdAt := now.Add(time.Duration(-i*24) * time.Hour)
		post := &model.Post{
			UserID:    user.ID,
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
