package errs

import "net/http"

type MessageErr interface {
	StatusCode() int
	Error() string
}

type ErrorData struct {
	ResponseCode int    `json:"statusCode"`
	ErrMessage   string `json:"message"`
}

func (e *ErrorData) StatusCode() int {
	return e.ResponseCode
}

func (e *ErrorData) Error() string {
	return e.ErrMessage
}

func NewUnauthorizedError(message string) MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusForbidden,
		ErrMessage:   message,
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusUnauthorized,
		ErrMessage:   message,
	}
}

func NewConflictError(message string) MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusConflict,
		ErrMessage:   message,
	}
}

func NewNotFoundError(message string) MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusNotFound,
		ErrMessage:   message,
	}
}

func NewBadRequest(message string) MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusBadRequest,
		ErrMessage:   message,
	}
}

func NewInternalServerError() MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusInternalServerError,
		ErrMessage:   "Something went wrong",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusUnprocessableEntity,
		ErrMessage:   message,
	}
}

func NewTimeOutError() MessageErr {
	return &ErrorData{
		ResponseCode: http.StatusRequestTimeout,
		ErrMessage:   "The request took too long to process. Please try again.",
	}
}
