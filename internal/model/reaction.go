package model

import "time"

type Reaction struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PostID    string    `json:"post_id" gorm:"type:uuid;not null;index:idx_post_user,unique" validate:"required"`
	Post      Post      `json:"-" gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE"`
	UserID    string    `json:"user_id" gorm:"type:uuid;not null;index:idx_post_user,unique" validate:"required"`
	User      User      `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Type      string    `json:"type" gorm:"type:text;not null;check:type IN ('like','dislike')" validate:"required,oneof=like dislike"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
