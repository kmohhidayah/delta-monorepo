package user

import (
	"context"
	"delta-monorepo/auth-app/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) error
	FindUserByPhoneNumber(ctx context.Context, phone string)(*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
  return &repository{db:db}
}

func (r *repository) CreateUser(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *repository) FindUserByPhoneNumber(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	return &user, err
}
