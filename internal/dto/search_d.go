package dto

import (
	"time"
)

type SearchRequest struct {
	Query string `json:"query" validate:"required,min=1,max=100"`
	Page  int    `json:"page" validate:"min=1"`
	Limit int    `json:"limit" validate:"min=1,max=50"`
}

type SearchResponse struct {
	Query       string       `json:"query"`
	CurrentPage int          `json:"current_page"`
	Users       []SearchUser `json:"users"`
}

type SearchUser struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	FollowerCount  int       `json:"follower_count"`
	FollowingCount int       `json:"following_count"`
	IsFollowing    bool      `json:"is_following"` // Whether current user follows this user
	JoinedAt       time.Time `json:"joined_at"`
}


