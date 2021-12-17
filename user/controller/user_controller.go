package controller

import (
	"BazarMessenger/common"
	"BazarMessenger/domain"
	"BazarMessenger/utility"
	"context"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(e *fiber.App, us domain.UserUsecase) {
	handler := &UserController{
		UserUsecase: us,
	}

	e.Get("/users/:username", handler.GetByUsername)
	//e.POST("/articles", handler.Store)
	//e.GET("/articles/:id", handler.GetByID)
	//e.DELETE("/articles/:id", handler.Delete)
}

func (this *UserController) GetByUsername(c *fiber.Ctx) error {
	if len(c.Params("username")) > 0 {
		username, err := this.UserUsecase.GetByUsername(c.Params("username"), context.TODO())
		if err != nil {
			dto := common.BaseDto{Message: utility.NOT_FOUND, StatusCode: 404}
			marshal := utility.ObjectToJson(dto)
			return c.SendString(marshal)
		} else {
			newUserDto := common.NewUserDto(common.BaseDto{Message: utility.SUCCESS, StatusCode: 200}, *username)
			marshal := utility.ObjectToJson(newUserDto)
			return c.SendString(marshal)
		}
	} else {
		dto := common.BaseDto{Message: utility.FAIL, StatusCode: 400}
		marshal := utility.ObjectToJson(dto)
		return c.SendString(marshal)
	}
}
