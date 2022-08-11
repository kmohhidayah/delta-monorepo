package dto

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorReply struct {
	Error ErrorMessage `json:"error"`
}

func ReplyError(message string) ErrorReply {
	return ErrorReply{
		Error: ErrorMessage{
			Message: message,
		},
	}
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
