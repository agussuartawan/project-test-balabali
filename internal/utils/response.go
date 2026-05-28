package utils

import (
	apperrors "github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func Success(c echo.Context, code int, message string, data any) error {
	return c.JSON(code, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c echo.Context, err error) error {
	appErr, ok := err.(*apperrors.AppError)
	if !ok {
		appErr = apperrors.NewInternalServerError()
	}

	return c.JSON(appErr.Code, ErrorResponse{
		Success: false,
		Message: appErr.Message,
		Error:   appErr.Details,
	})
}