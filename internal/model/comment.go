package model

import "time"

type Comment struct {
	ID              string
	PostID          string
	UserID          string
	Content         string
	ParentCommentID *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
