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
type UserErrors struct {
	Err    bool   `json:"error"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

func NewUserDto(baseDto BaseDto, user domain.User) *UserDto {
	return &UserDto{BaseDto: baseDto, Data: user}
}
func NewUserDto2(baseDto BaseDto) *UserDto {
	return &UserDto{BaseDto: baseDto}
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
