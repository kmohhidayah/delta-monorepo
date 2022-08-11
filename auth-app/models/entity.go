package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
}

type User struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
