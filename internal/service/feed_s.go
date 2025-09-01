package service

import (
	"log"
	"soulvent/internal/model"
	"soulvent/internal/repository"
)


func GetUserFeed(userID string, pageNum int, limit int) ([]model.Post, error) {

	seenPostIDs,err := repository.GetSeenPostIDs(userID)
	if err != nil {
		log.Println("Error fetching seen posts:", err)
	}
	followingIDs, err := repository.GetFollowingIDs(userID)
	if err != nil {
		log.Println("Error fetching following IDs:", err)
		return nil, err
	}
	log.Println("followingIds: ", followingIDs)

	posts, err := repository.GetFeedPosts(followingIDs, seenPostIDs, pageNum, limit)
	if err != nil {
		log.Println("Error fetching feed posts:", err)
		return nil, err
	}
	//may add trending posts logic here
	return posts, nil
}