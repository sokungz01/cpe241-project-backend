package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type ServiceResponseController interface {
	GetAll(c *fiber.Ctx) error
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
