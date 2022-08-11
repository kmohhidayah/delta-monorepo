package user

import (
	"context"
	"delta-monorepo/auth-app/models"
	"delta-monorepo/auth-app/util"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, input RegisterUserInput) (*RegisterUserResponse, error)
	Login(ctx context.Context, input LoginInput) (*LoginResponse, error)
	Claim(ctx context.Context, tokenString string) (*TokenResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(ctx context.Context, input RegisterUserInput) (*RegisterUserResponse, error) {
	password := util.RandomPassword()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateUser(ctx, models.User{
		Model: models.Model{
			ID: uuid.New().String(),
		},
		Name:     input.Name,
		Phone:    input.Phone,
		Role:     input.Role,
		Password: string(hashedPass),
	})
	if err != nil {
		return nil, err
	}

	return &RegisterUserResponse{
		Name:     input.Name,
		Phone:    input.Phone,
		Role:     input.Role,
		Password: password,
	}, nil
}

func (s *service) Login(ctx context.Context, input LoginInput) (*LoginResponse, error) {
	user, err := s.repository.FindUserByPhoneNumber(ctx, input.Phone)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &TokenResponse{
		Name:  user.Name,
		Phone: user.Phone,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenString,
	}, nil
}

func (a *service) Claim(ctx context.Context, tokenString string) (*TokenResponse, error) {
	tokenStr := strings.Replace(tokenString, "Bearer ", "", -1)

	token, err := jwt.ParseWithClaims(tokenStr, &TokenResponse{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenResponse); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
