package service

import (
	"log"
	"soulvent/internal/model"
	"soulvent/internal/repository"
	"time"
)

func GetUserFeed(userID string, limit int) ([]model.Post, error) {

	seenPostIDs, err := repository.GetSeenPostIDs(userID)
	if err != nil {
		log.Println("Error fetching seen posts:", err)
	}
	log.Println("seen post ids of user: ", userID, "PostIDs: ", seenPostIDs)
	followingIDs, err := repository.GetFollowingIDs(userID)
	if err != nil {
		log.Println("Error fetching following IDs:", err)
		return nil, err
	}
	log.Println("followingIds: ", followingIDs)

	posts, err := repository.GetFeedPosts(followingIDs, seenPostIDs, limit)
	if err != nil {
		log.Println("Error fetching feed posts:", err)
		return nil, err
	}
	//may add trending posts logic here
	//....
	postIDs := make([]string, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	go MarkPostsSeen(userID, postIDs)
	return posts, nil
}

func MarkPostsSeen(userID string, postIDs []string) error {
	for _, postID := range postIDs {
		if err := repository.MarkPostSeen(userID, postID); err != nil {
			log.Println("failed to mart post viwed . postID :", postID, err)
			continue
		}
	}
	return nil
}

func ClearOldSeenPost(userId string, cutoffDate string) error {
	cutoffTime, err := time.Parse("2006-01-02 15:04:05", cutoffDate)
	if err != nil {
		return err
	}
	cutoffTimeUnix := cutoffTime.Unix()

	if err := repository.ClearOldSeenPost(userId, cutoffTimeUnix); err != nil {
		return err
	}
	return nil
}
