package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type serviceRequestController struct {
	serviceRequestUsecase domain.ServiceRequestUsecase
}

type ServiceRequestController interface {
	GetAllServiceRequest(c *fiber.Ctx) error
	GetServiceRequest(c *fiber.Ctx) error
	CreateServiceRequest(c *fiber.Ctx) error
}

func NewServiceRequestController(serviceRequestUsecase domain.ServiceRequestUsecase) ServiceRequestController {
	return &serviceRequestController{
		serviceRequestUsecase: serviceRequestUsecase,
	}
}

func (srq *serviceRequestController) GetAllServiceRequest(c *fiber.Ctx) error {
	response, err := srq.serviceRequestUsecase.GetAllServiceRequest()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (srq *serviceRequestController) GetServiceRequest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err := srq.serviceRequestUsecase.GetServiceRequest(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (srq *serviceRequestController) CreateServiceRequest(c *fiber.Ctx) error {
	input := new(domain.ServiceRequest)
	parseErr := c.BodyParser(input)
	if parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseErr.Error())
	}
	srq.serviceRequestUsecase.CreateServiceRequest(input)
	return c.SendStatus(fiber.StatusOK)
}
