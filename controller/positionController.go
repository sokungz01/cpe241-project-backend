package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type positionController interface {
	CreatePosition(c *fiber.Ctx) error
	GetByPositionName(c *fiber.Ctx) error
	GetByPositionID(c *fiber.Ctx) error
	UpdatePosition(c *fiber.Ctx) error
	DeletePosition(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type positionUsecase struct {
	posusecase domain.PositionUsecase
}

func NewPositionController(posusecase domain.PositionRepository) positionController {
	return &positionUsecase{posusecase: posusecase}
}

func (usecase *positionUsecase) CreatePosition(c *fiber.Ctx) error {
	newPosition := new(domain.Position)
	if err := c.BodyParser(newPosition); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)

	}
	if err := usecase.posusecase.Create(newPosition); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

func (usecase *positionUsecase) GetByPositionName(c *fiber.Ctx) error {
	Parse := new(domain.Position)
	if err := c.BodyParser(Parse); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	response, err := usecase.posusecase.FindByPositionName(Parse.PositionName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (usecase *positionUsecase) GetByPositionID(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(convErr.Error())
	}
	response, getErr := usecase.posusecase.GetPositionByID(id)
	if getErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(getErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (usecase *positionUsecase) UpdatePosition(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	newData := new(domain.Position)
	if convErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(convErr.Error())
	}
	if parseErr := c.BodyParser(newData); parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseErr.Error())
	}

	if updateErr := usecase.posusecase.UpdatePosition(id, newData); updateErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(updateErr.Error())
	}

	response, getErr := usecase.posusecase.GetPositionByID(id)
	if getErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(getErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (usecase *positionUsecase) DeletePosition(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	deleteErr := usecase.posusecase.DeletePosition(id)
	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(deleteErr.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (usecase *positionUsecase) GetAll(c *fiber.Ctx) error {
	response, err := usecase.posusecase.GetAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
