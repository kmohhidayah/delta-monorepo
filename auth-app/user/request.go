package user

import validation "github.com/go-ozzo/ozzo-validation"

type RegisterUserInput struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

type LoginInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *RegisterUserInput) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Phone, validation.Required),
		validation.Field(&r.Role, validation.Required),
	)
}

func (l *LoginInput) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Phone, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}
