package service

import(
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/repository"
)

func CreateFollower(followerReq *dto.CreateFollowerRequest) error {
	follower := model.Follower{
		UserID:     followerReq.UserID,
		FollowerID: followerReq.FollowerID,
	}
	if err := repository.CreateFollower(&follower); err != nil {
		return err
	}
	return nil
}

func GetFollowers(userID string, followerID string) ([]model.Follower, error) {
	followers, err := repository.GetFollowers(userID, followerID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}