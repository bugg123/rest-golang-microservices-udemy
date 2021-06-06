package errs

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(message, http.StatusNotFound)
}
func NewUnexpectedError(message string) *AppError {
	return NewAppError(message, http.StatusInternalServerError)
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
