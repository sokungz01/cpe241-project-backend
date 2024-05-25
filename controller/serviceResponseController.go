package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ServiceResponseController interface {
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	DeleteResponse(c *fiber.Ctx) error
	CreateResponse(c *fiber.Ctx) error
}

type serviceResponseController struct {
	usecase domain.ServiceResponseUsecase
}

func NewServiceResponseController(usecase domain.ServiceResponseUsecase) serviceResponseController {
	return serviceResponseController{usecase: usecase}
}

func (u *serviceResponseController) GetAll(c *fiber.Ctx) error {
	response, err := u.usecase.GetAllResponse()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *serviceResponseController) GetOne(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(convErr.Error())
	}
	response, getErr := u.usecase.GetResponse(id)
	if getErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(getErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *serviceResponseController) DeleteResponse(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(convErr.Error())
	}
	deleteErr := u.usecase.DeleteResponse(id)
	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(deleteErr.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (u *serviceResponseController) CreateResponse(c *fiber.Ctx) error {
	input := new(domain.ServiceResponse)
	parseErr := c.BodyParser(input)
	if parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseErr.Error())
	}
	createErr := u.usecase.CreateServiceResponse(input)
	if createErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(createErr.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
