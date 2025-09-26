package repository

import (
	"fmt"
	"soulvent/internal/db"
	"soulvent/internal/dto"
	"soulvent/internal/model"
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
			ID:             user.ID,
			Username:       user.Username,
			IsFollowing:    isFollowing(userId, user.ID),
			JoinedAt:       user.CreatedAt,
		}
		results = append(results, userResult)
	}
	return results, nil
}

func AddSearch(addSearchReq *model.Search)error{
	if err := db.PgDb.Create(addSearchReq).Error; err != nil {
		return err
	}
	return nil
}

func GetSearches(userID string, limit int)([]*model.Search, error){
	var searches []*model.Search
	if err := db.PgDb.Order("created_at DESC").Limit(5).Find(&searches).Error; err != nil{
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
		Where("user_id = ? AND FollowerID = ?", targetUserID,currentUserID).
		Count(&count)
	return count > 0
}
