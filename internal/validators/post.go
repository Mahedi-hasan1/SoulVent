package validators

import (
	"errors"
	"soulvent/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateCreatePost(postReq *dto.CreatePostRequest, userID string) error {
	validate := validator.New()
	 if err := validate.Struct(postReq); err != nil {
        return err
    }
	if postReq == nil {
		return errors.New("request is nil")
	}

	if userID == "" {
		return errors.New("user_id is required")
	}

	if postReq.Content == "" &&  len(postReq.ImageURLs) == 0 {
		return errors.New("at least one of 'content' or 'image_urls' must be provided")
	}

	for _, url := range postReq.ImageURLs {
		if len(url) > 500 {
			return errors.New("each image URL must not exceed 500 characters")
		}
	}
	return nil
}

func ValidateGetUserPosts(username string, limit int) error {

	if username == "" {
		return errors.New("username is required")
	}
	if limit <= 0 || limit > 100 {
		return errors.New("limit must be a positive integer between 1 and 100")
	}
	return nil
}