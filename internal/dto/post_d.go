package dto

import "time"

type CreatePostRequest struct {
	//UserID    string   `json:"user_id" validate:"required,uuid4"`
	Content   string   `json:"content" validate:"max=10000"`
	ImageURLs []string `json:"image_urls,omitempty" validate:"omitempty,dive,url,max=100"` // max 100 images, each a valid URL
}

type UserPostResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Content      string    `json:"content"`
	ImageURLs    []string  `json:"image_urls" gorm:"serializer:json;type:json"`
	LikeCount    int       `json:"like_count"`
	DislikeCount int       `json:"dislike_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
