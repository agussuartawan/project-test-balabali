package utils

import (
	"strconv"

	apperrors "github.com/agussuartawan/project-test-balabali/internal/errors"
)

func ParseID(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 64); if err != nil {
		return 0, apperrors.NewBadRequestError("invalid id format", nil)
	}

	return uint(id), nil
}