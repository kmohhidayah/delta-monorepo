package user

import (
	"delta-monorepo/auth-app/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByPhoneNumber(phone string) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindUserByPhoneNumber(phone string) (models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
