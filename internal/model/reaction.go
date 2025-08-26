package model

type Reaction struct {
	ID        string
	PostID    string
	UserID    string
	Type      string // e.g. love, dislike etc.
	CreatedAt int64
}
