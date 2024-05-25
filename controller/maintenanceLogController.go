package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type MaintenanceLogController interface {
	GetAllmaintenanceLog(c *fiber.Ctx) error
}

type maintenanceLogcontroller struct {
	usecase domain.MaintenanceLogUsecase
}

func NewMaintenanceLogController(usecase domain.MaintenanceLogUsecase) MaintenanceLogController {
	return &maintenanceLogcontroller{usecase: usecase}
}

func (u *maintenanceLogcontroller) GetAllmaintenanceLog(c *fiber.Ctx) error {
	response, err := u.usecase.GetAllmaintenanceLog()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
