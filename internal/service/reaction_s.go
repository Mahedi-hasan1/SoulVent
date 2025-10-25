package service

import (
	"soulvent/internal/model"
	"soulvent/internal/repository"
)

func AddReaction(reactionAddReq *model.Reaction) error {
	return repository.AddReaction(reactionAddReq)
}

func RemoveReaction(postID, userID string) error {
	return repository.RemoveReaction(postID,userID)
}