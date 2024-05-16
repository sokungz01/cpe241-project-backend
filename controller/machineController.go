package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type MachineController interface {
	CreateMachine(c *fiber.Ctx) error
	GetAllMachine(c *fiber.Ctx) error
	GetMachineByID(c *fiber.Ctx) error
}

type machineController struct {
	usecase domain.MachineUsecase
}

func NewMachineController(usecase domain.MachineUsecase) MachineController {
	return &machineController{usecase: usecase}
}

func (m *machineController) CreateMachine(c *fiber.Ctx) error {
	newMachine := new(domain.Machine)
	var err error
	//Parsing JSON
	if err := c.BodyParser(newMachine); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	//Creat Machine
	_, err = m.usecase.CreateMachine(newMachine)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(newMachine)
}

func (m *machineController) GetAllMachine(c *fiber.Ctx) error {
	response, err := m.usecase.GetAllMachine()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (m *machineController) GetMachineByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, errorResponse := m.usecase.GetMachineByID(id)
	if errorResponse != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(errorResponse.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
