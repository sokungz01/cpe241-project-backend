package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type MaintenanceStatusController interface {
	GetAll(c *fiber.Ctx) error
}

type maintenanceStatusController struct {
	usecase domain.MaintenanceStatusUsecase
}

func NewMaintenanceStatuscontroller(usecase domain.MaintenanceStatusUsecase) MaintenanceStatusController {
	return &maintenanceStatusController{usecase: usecase}
}

func (u *maintenanceStatusController) GetAll(c *fiber.Ctx) error {
	response, err := u.usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}


