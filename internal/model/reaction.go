package model

import "time"

type Reaction struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PostID    string    `json:"post_id" gorm:"type:uuid;not null;index:idx_post_user,unique"`
	Post      Post      `json:"post" gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE"`
	UserID    string    `json:"user_id" gorm:"type:uuid;not null;index:idx_post_user,unique"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Type      string    `json:"type" gorm:"type:text;not null;check:type IN ('like','dislike')"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
