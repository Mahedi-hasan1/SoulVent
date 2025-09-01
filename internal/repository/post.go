package repository

import (
	"log"
	"soulvent/internal/db"
	"soulvent/internal/model"
)

func CreatePost(post *model.Post) error {
	// Logic to save post to the database
	if err := db.PgDb.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func GetPosts(postID string, userID string) ([]model.Post, error) {
	var posts []model.Post
	query := db.PgDb.Model(&model.Post{})
	if postID != "" {
		query = query.Where("id = ?", postID)
	}
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetFeedPosts(followingIDs []string, seenPostIDs []string, pageNum int, limit int) ([]model.Post, error) {
	var posts []model.Post
	//offset := (pageNum-1)*limit
	// I think dont need offset as we getting always unseen post . frontend request limit amount of next posts
	// then this api will return those amount of posts.
	log.Println("seen post id ", seenPostIDs)
	query := db.PgDb.Model(&model.Post{}).
		Where("user_id IN ?", followingIDs).
		Order("hot_score DESC").
		Order("created_at DESC").
		Limit(limit)

	if len(seenPostIDs) > 0 {
		query = query.Where("id NOT IN ?", seenPostIDs)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
