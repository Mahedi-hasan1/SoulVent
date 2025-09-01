package dto

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Gender   string `json:"gender" validate:"required,oneof=male female other"`
	City     string `json:"city" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=8"` // Plain password in request
}