package validators

import (
	"errors"
	"soulvent/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateCreatePost(postReq *dto.CreatePostRequest) error {
	validate := validator.New()
	 if err := validate.Struct(postReq); err != nil {
        return err
    }
	if postReq == nil {
		return errors.New("request is nil")
	}

	if postReq.UserID == "" {
		return errors.New("user_id is required")
	}

	if postReq.Content == "" && (postReq.ImageURLs == nil || len(postReq.ImageURLs) == 0) {
		return errors.New("at least one of 'content' or 'image_urls' must be provided")
	}

	for _, url := range postReq.ImageURLs {
		if len(url) > 500 {
			return errors.New("each image URL must not exceed 500 characters")
		}
	}
	return nil
}