package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type analysisController struct {
	repo domain.AnalysisRepository
}

func NewAnalysisController(repo domain.AnalysisRepository) domain.InventoryAnalysisController {
	return &analysisController{repo: repo}
}

func (ac *analysisController) GetInventoryAnanlysis(c *fiber.Ctx) error {
	response, err := ac.repo.GetInventoryAnanlysis()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ac *analysisController) GetMachineTypeErrorAnalysis(c *fiber.Ctx) error {
	response, err := ac.repo.GetMachineTypeErrorAnalysis()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ac *analysisController) GetEmployeeEngagementAnalysis(c *fiber.Ctx) error {
	response, err := ac.repo.GetEmployeeEngagementAnalysis()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ac *analysisController) GetMaintenanceCostAnalysis(c *fiber.Ctx) error {
	response, err := ac.repo.GetMaintenanceCostAnalysis()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
