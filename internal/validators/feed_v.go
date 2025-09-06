package validators

import (
	"errors"
)

func ValidateGetUserFeed(userID string, limit int) error {
	if userID == "" {
		return errors.New("user_id is required")
	}
	if limit <= 0 || limit > 100 {
		return errors.New("limit must be a positive integer between 1 and 100")
	}
	return nil
}


func ValidateClearOldSeenPosts(userID, date string) error {
	if userID == "" {
		return errors.New("user_id is required")
	}

	return nil
}