package user

import "github.com/golang-jwt/jwt"

type RegisterUserResponse struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type TokenResponse struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
