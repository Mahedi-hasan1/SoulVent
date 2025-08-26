package model

type Post struct {
	ID        string
	UserID    string
	Content   string
	ImageURLs  []string
	CreatedAt int64
	ReactionCount int
	CommentCount  int
	HotScore      float64
}
