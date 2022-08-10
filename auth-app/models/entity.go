package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
