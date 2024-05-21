package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ItemLogController interface {
	GetAll(ctx *fiber.Ctx) error
}

type itemLogController struct {
	usecase domain.ItemLogUsecase
}

func NewItemLogController(usecase domain.ItemLogUsecase) ItemLogController {
	return &itemLogController{usecase: usecase}
}

func (c *itemLogController) GetAll(ctx *fiber.Ctx) error {
	response, err := c.usecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
