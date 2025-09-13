package dto

type CreateFollowerRequest struct {
	UserID     string `json:"user_id" binding:"required,uuid"`
	FollowerID string `json:"-"`
}
