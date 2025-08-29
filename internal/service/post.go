package service

import (
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/repository"
)

func CreatePost(postCreateReq *dto.CreatePostRequest) error {
	post := &model.Post{
		UserID:  postCreateReq.UserID,
		Content: postCreateReq.Content,
		ImageURLs: postCreateReq.ImageURLs,
	}
	return repository.CreatePost(post); 
}
