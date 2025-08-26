package repository

// import (
// 	"SoulVent/internal/model"

// 	"gorm.io/gorm"
// )

// type UserRepository interface {
// 	Create(user *model.User) error
// 	GetByID(id string) (*model.User, error)
// }

// type userRepo struct {
// 	DB *gorm.DB
// }

// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userRepo{DB: db}
// }

// func (r *userRepo) Create(user *model.User) error {
// 	return r.DB.Create(user).Error
// }
// func (r *userRepo) GetByID(id string) (*model.User, error) {
// 	var user model.User
// 	err := r.DB.First(&user, "id = ?", id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
