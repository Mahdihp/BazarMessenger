package common

import (
	"BazarMessenger/domain"
	"BazarMessenger/utility"
	"net/http"
)

type BaseDto struct {
	StatusCode int         `json:"statusCode"`
	Message    utility.Key `json:"message"`
}

type UserDto struct {
	BaseDto
	Data domain.User `json:"data"`
}

func NewUserDto(baseDto BaseDto, user domain.User) *UserDto {
	return &UserDto{BaseDto: baseDto, Data: user}
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
