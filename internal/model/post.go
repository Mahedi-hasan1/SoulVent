package model

type Post struct {
	ID        string
	UserID    string
	Content   string
	ImageURL  string
	CreatedAt int64
	Trending  bool
}
