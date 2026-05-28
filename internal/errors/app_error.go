package errors

import "net/http"

type AppError struct {
	Code    int
	Message string
	Details any
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string, details any) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		Details: details,
	}
}

func NewUnprocessableEntityError(message string, details any) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
		Details: details,
	}
}

func NewNotFoundError(message string, details any) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
		Details: details,
	}
}

func NewInternalServerError() *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Details: nil,
	}
}