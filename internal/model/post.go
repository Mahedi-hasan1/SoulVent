package model

import "time"

type Post struct {
	ID            string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID        string    `json:"user_id" gorm:"type:uuid;not null;index"`
	User          User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Content       string    `json:"content" gorm:"type:text;not null"`
	ImageURLs     []string  `json:"image_urls" gorm:"serializer:json;type:json"`
	ReactionCount int       `json:"reaction_count" gorm:"default:0"`
	CommentCount  int       `json:"comment_count" gorm:"default:0"`
	HotScore      float64   `json:"hot_score" gorm:"type:decimal(10,6);default:0"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
