package repository

import (
	"soulvent/internal/db"
	"soulvent/internal/model"

	"gorm.io/gorm"
)

func AddReaction(reactionAddReq *model.Reaction) error {
	var existing model.Reaction
	postID := reactionAddReq.PostID
	err := db.PgDb.Where("post_id = ? and user_id = ?", postID, reactionAddReq.UserID).
		First(&existing).Error
	//new reaction add, incresing reaction count in post and it should be transactional
	if err == gorm.ErrRecordNotFound {
		return db.PgDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(reactionAddReq).Error; err != nil {
			return err
		} else {
			return AdjustLikeOrDislikeCount(tx, "+",postID, reactionAddReq.Type)
		}
	})
	}else if err != nil {
		return err
	}
	
	if existing.Type == reactionAddReq.Type{
		return nil
	}
	//updating reaction
	return db.PgDb.Transaction(
		func(tx *gorm.DB) error {
			if err := AdjustLikeOrDislikeCount(tx, "-", postID, existing.Type); err != nil{
				return err
			}
			if err := AdjustLikeOrDislikeCount(tx, "+",postID, reactionAddReq.Type); err != nil{
				return err
			}
			existing.Type=reactionAddReq.Type
			return tx.Save(&existing).Error
		})
}

func RemoveReaction(postID, userID string) error {
	var reaction model.Reaction
	return db.PgDb.Transaction(
		func(tx *gorm.DB) error {
			err:= tx.Where("post_id = ? and user_id = ?", postID, userID).
			First(&reaction).Error
			if err != nil {
			if err == gorm.ErrRecordNotFound{
				return nil
			}else {
				return err
			}}

			if err := AdjustLikeOrDislikeCount(tx, "-",postID, reaction.Type ); err!= nil{
				return err
			}
			return tx.Where("post_id = ? and user_id = ?", postID, userID).
				Delete(&model.Reaction{}).Error
		})		
}

func AdjustLikeOrDislikeCount(tx *gorm.DB, operator,postID,reactType string) error {
	switch reactType {
	case "like":
		return tx.Model(&model.Post{}).Where("id = ?", postID).
			UpdateColumn("like_count", gorm.Expr("like_count "+operator+" ?", 1)).Error
	case "dislike":
		return tx.Model(&model.Post{}).Where("id = ?", postID).
			UpdateColumn("dislike_count", gorm.Expr("dislike_count "+operator+" ?", 1)).Error
	default:	
		return nil
	}
}
