package validators

import (
	"errors"
	"soulvent/internal/dto"
	"soulvent/internal/util"

	"github.com/go-playground/validator/v10"
)
func ValidateCreateUser(userReq *dto.CreateUserRequest) error {
	// Implement user validation logic
	validate := validator.New()
	 if err := validate.Struct(&userReq); err != nil {
        customErrorMsg := util.GetCustomValidationMessage(err)
        return errors.New(customErrorMsg)
    }
	return nil
}