package service

import (
	"log"
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

func GetPosts(postID string, userID string) ([]model.Post, error) {
	return repository.GetPosts(postID, userID); 
}
func BulkCreatePost(postsCreateReq *[]dto.CreatePostRequest) error {
	for _, postCreateReq := range *postsCreateReq{
		post := &model.Post{
			UserID:  postCreateReq.UserID,
			Content: postCreateReq.Content,
			ImageURLs: postCreateReq.ImageURLs,
		}
		if err := repository.CreatePost(post); err != nil{
			log.Panicln("post not created. details: ", post)
		}else{
			log.Println("post created: details", post.Content)
		}
	} 
	return nil
}