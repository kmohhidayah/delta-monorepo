package errors

import "delta-monorepo/auth-app/dto"

type Error struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Unknown error"`
}

func (app *Error) Error() string {
	return app.Message
}

func (app *Error) ToReply() *dto.ErrorReply {
	return &dto.ErrorReply{
		Error: dto.ErrorMessage{
			Message: app.Message,
		},
	}
}

// NewError creates a new Error with an optional message
func NewError(code int, message string) *Error {
	e := &Error{
		Code:    code,
		Message: message,
	}
	return e
}
