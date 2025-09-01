package model

type Follower struct {
	ID         string  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID     string  `json:"user_id" gorm:"type:uuid;not null;index;uniqueIndex:uniq_follower" `
	FollowerID string  `json:"follower_id" gorm:"type:uuid;not null;index;uniqueIndex:uniq_follower"`

	User	   User    `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Follower		User  `json:"follower" gorm:"foreignKey:FollowerID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt  string  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}