package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ItemCategoryController interface {
	CreateItemCategory(ctx *fiber.Ctx) error
	GetAllItemCategory(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	UpdateItemCategory(ctx *fiber.Ctx) error
	DeleteItemCategory(ctx *fiber.Ctx) error
}

type itemCategoryController struct {
	usecase domain.ItemCategoryUseCase
}

func NewItemCategoryController(usecase domain.ItemCategoryUseCase) ItemCategoryController {
	return &itemCategoryController{usecase: usecase}
}

func (c *itemCategoryController) CreateItemCategory(ctx *fiber.Ctx) error {
	newItemCategory := new(domain.ItemCategory)
	if err := ctx.BodyParser(newItemCategory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err := c.usecase.CreateItemCategory(newItemCategory)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *itemCategoryController) GetAllItemCategory(ctx *fiber.Ctx) error {
	response, err := c.usecase.GetAllItemCategory()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *itemCategoryController) FindByID(ctx *fiber.Ctx) error {
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

func (c *itemCategoryController) UpdateItemCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newItemCategory := new(domain.ItemCategory)
	if err := ctx.BodyParser(newItemCategory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response, err := c.usecase.UpdateItemCategory(id, newItemCategory)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *itemCategoryController) DeleteItemCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = c.usecase.DeleteItemCategory(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Delete Item category successfully")
}
