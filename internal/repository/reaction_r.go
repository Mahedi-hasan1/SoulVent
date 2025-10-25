package repository

import (
	"soulvent/internal/db"
	"soulvent/internal/model"

	"gorm.io/gorm"
)

func AddReaction(reactionAddReq *model.Reaction) error {
	var existing model.Reaction
	err := db.PgDb.Where("post_id = ? and user_id = ?", reactionAddReq.PostID, reactionAddReq.UserID).
		First(&existing).Error
	if err == gorm.ErrRecordNotFound {
		return db.PgDb.Create(reactionAddReq).Error
	}
	if err!= nil {
		return err
	}
	existing.Type = reactionAddReq.Type
	return db.PgDb.Save(&existing).Error
}

func RemoveReaction(postID, userID string) error {
	return db.PgDb.Where("post_id = ? and user_id = ?", postID, userID).
		Delete(&model.Reaction{}).Error
}
