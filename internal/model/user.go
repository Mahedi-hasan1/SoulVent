package model

import "time"

type User struct {
	ID           string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username     string    `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Gender       string    `json:"gender" gorm:"type:varchar(10);check:gender IN ('male','female','other')"`
	City         string    `json:"city" gorm:"not null;size:50"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null;size:100"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
