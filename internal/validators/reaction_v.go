package validators

import (
	"errors"
	"soulvent/internal/model"

	"github.com/go-playground/validator/v10"
)

func ValidateAddReaction(addReactionReq *model.Reaction) error{
	if addReactionReq.PostID =="" || addReactionReq.UserID ==""{
		return errors.New("post_id and user_id is required")	
	}
	validate := validator.New();
	if err := validate.Struct(addReactionReq); err != nil {
		return err
	}
	return nil
} 