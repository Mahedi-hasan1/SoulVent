package validators
import (
	"errors"
	"soulvent/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateFollower(followerReq *dto.CreateFollowerRequest) error {
	var validate = validator.New()
	if err := validate.Struct(followerReq); err != nil {
		return err
	}
	if followerReq.UserID == followerReq.FollowerID {
		return errors.New("a user cannot follow themselves")
	}

	return nil
}

func ValidateGetFollowers(userID string, followerID string) error {
	if userID == "" && followerID == "" {
		return errors.New("at least one of 'user_id' or 'follower_id' must be provided")
	}
	return nil
}