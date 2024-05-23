package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ItemController interface {
	CreateItem(ctx *fiber.Ctx) error
	GetAllItem(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	UpdateItem(ctx *fiber.Ctx) error
	DeleteItem(ctx *fiber.Ctx) error
}

type itemController struct {
	usecase domain.ItemUseCase
}

func NewItemController(usecase domain.ItemUseCase) ItemController {
	return &itemController{usecase: usecase}
}

func (c *itemController) CreateItem(ctx *fiber.Ctx) error {
	newItem := new(domain.Item)
	if err := ctx.BodyParser(newItem); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err := c.usecase.CreateItem(newItem)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (c *itemController) GetAllItem(ctx *fiber.Ctx) error {
	response, err := c.usecase.GetAllItem()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (c *itemController) FindByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err := c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (c *itemController) UpdateItem(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newItem := new(domain.Item)
	if err := ctx.BodyParser(newItem); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response, err := c.usecase.UpdateItem(id, newItem)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (c *itemController) DeleteItem(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = c.usecase.DeleteItem(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Delete Item successfully")
}
