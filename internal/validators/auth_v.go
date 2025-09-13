package validators

import (
	"errors"
	"soulvent/internal/dto"
)

func ValidateLogIn(logINReq *dto.LogInReq)error {
	if  logINReq.UsernameOrEmail == "" {
		return errors.New("username or email is required")
	}
	return nil
}