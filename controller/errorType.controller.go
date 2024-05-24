package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ErrorTypeController interface {
	CreateErrorType(ctx *fiber.Ctx) error
	GetAllErrorType(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	UpdateErrorType(ctx *fiber.Ctx) error
	DeleteErrorType(ctx *fiber.Ctx) error
}

type errorTypeController struct {
	usecase domain.ErrorTypeUseCase
}

func NewErrorTypeController(usecase domain.ErrorTypeUseCase) ErrorTypeController {
	return &errorTypeController{usecase: usecase}
}

func (c *errorTypeController) CreateErrorType(ctx *fiber.Ctx) error {
	newErrorType := new(domain.ErrorType)
	if err := ctx.BodyParser(newErrorType); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err := c.usecase.CreateErrorType(newErrorType)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *errorTypeController) GetAllErrorType(ctx *fiber.Ctx) error {
	response, err := c.usecase.GetAllErrorType()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *errorTypeController) FindByID(ctx *fiber.Ctx) error {
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

func (c *errorTypeController) UpdateErrorType(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newErrorType := new(domain.ErrorType)
	if err := ctx.BodyParser(newErrorType); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response, err := c.usecase.UpdateErrorType(id, newErrorType)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *errorTypeController) DeleteErrorType(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_, err = c.usecase.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = c.usecase.DeleteErrorType(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Delete Item category successfully")
}
