package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ItemLogController interface {
	GetAll(ctx *fiber.Ctx) error
	CreateItemLog(ctx *fiber.Ctx) error
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

func (c *itemLogController) CreateItemLog(ctx *fiber.Ctx) error {
	newItemLog := new(domain.ItemLog)
	if err := ctx.BodyParser(newItemLog); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if newItemLog.ItemID == 0 || newItemLog.StaffID == 0 || newItemLog.ItemQty == 0 {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	response, err := c.usecase.CreateItemLog(newItemLog)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
