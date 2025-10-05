package repository

import (
	"errors"
	"fmt"
	"soulvent/internal/db"
	"soulvent/internal/dto"
	"soulvent/internal/model"

	"gorm.io/gorm"
)

func SearchUsers(query, userId string, page, limit int) ([]dto.SearchUser, error) {
	offset := (page - 1) * limit
	var users []model.User

	// Search users and order by relevance (username match first, then by creation date)
	err := db.PgDb.Model(&model.User{}).
		Where("username ILIKE ? AND id != ?", "%"+query+"%", userId).
		Order(fmt.Sprintf("CASE WHEN username ILIKE '%s%%' THEN 1 ELSE 2 END", query)).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	// Convert to UserResult DTOs
	var results []dto.SearchUser
	for _, user := range users {
		userResult := dto.SearchUser{
			ID:          user.ID,
			Username:    user.Username,
			IsFollowing: isFollowing(userId, user.ID),
			JoinedAt:    user.CreatedAt,
		}
		results = append(results, userResult)
	}
	return results, nil
}

func AddSearch(searchReq *model.Search) error {
	var existingSearch model.Search
	err := db.PgDb.Where("user_id = ? AND query = ?", searchReq.UserID, searchReq.Query).First(&existingSearch).Error
	if err == nil {
		if err := db.PgDb.Model(&existingSearch).
			Updates(map[string]interface{}{}).Error; err != nil {
			return errors.New("failed to update updated_at field: " + err.Error())
		}
		return nil
	}

	if err != gorm.ErrRecordNotFound {
		return err
	}

	if err := db.PgDb.Create(searchReq).Error; err != nil {
		return errors.New("failed to add search : " + err.Error())
	}
	return nil
}

func GetSearches(userID string, limit int) ([]*model.Search, error) {
	var searches []*model.Search
	if err := db.PgDb.Order("updated_at DESC").Limit(5).Find(&searches).Error; err != nil {
		return nil, err
	}
	return searches, nil
}

func isFollowing(currentUserID, targetUserID string) bool {
	if currentUserID == targetUserID {
		return false
	}
	var count int64
	db.PgDb.Model(&model.Follower{}).
		Where("user_id = ? AND follower_id = ?", targetUserID, currentUserID).
		Count(&count)
	return count > 0
}
