package model

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Content   string
	CreatedAt int64
}
