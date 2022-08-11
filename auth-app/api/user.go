package api

import (
	"delta-monorepo/auth-app/dto"
	"delta-monorepo/auth-app/errors"
	"delta-monorepo/auth-app/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
	Claims(c *gin.Context)
}

type userAPI struct {
	userService user.Service
}

func NewUserAPI(userService user.Service) UserAPI {
  return &userAPI{userService: userService}
}

func (u *userAPI) Login(c *gin.Context) {
	var request user.LoginInput

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := u.userService.Login(c, request)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
	}

	res := dto.SuccessResponse{
		Message: "success",
		Data:    response,
	}

	c.JSON(http.StatusOK, res)
}

func (u *userAPI) CreateUser(c *gin.Context) {
	var request user.RegisterUserInput

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := u.userService.CreateUser(c, request)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := dto.SuccessResponse{
		Message: "success",
		Data:    response,
	}

	c.JSON(http.StatusOK, res)
}

func (u *userAPI) Claims(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Error(errors.NewError(http.StatusBadRequest, "token is required"))
		return
	}

	response, err := u.userService.Claim(c, token)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := dto.SuccessResponse{
		Message: "success",
		Data:    response,
	}

	c.JSON(http.StatusOK, res)
}
