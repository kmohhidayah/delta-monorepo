package user

import "github.com/golang-jwt/jwt"

type RegisterUserInput struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

type LoginInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenInput struct {
  Name string `json:"name"`
  Phone string `json:"phone"`
  Role string `json:"role"`
  jwt.StandardClaims
}
