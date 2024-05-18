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
	GetMachineByName(c *fiber.Ctx) error
	UpdateMachineData(c *fiber.Ctx) error
	DeleteMachine(c *fiber.Ctx) error
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

func (m *machineController) GetMachineByName(c *fiber.Ctx) error {
	input := new(domain.Machine)
	parseError := c.BodyParser(input)
	if parseError != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseError.Error())
	}
	response, getErr := m.usecase.GetMachineByName(input.MachineName)
	if getErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(getErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (m *machineController) UpdateMachineData(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	newMachineData := new(domain.Machine)

	if parseErr := c.BodyParser(newMachineData); parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseErr.Error())
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, updateError := m.usecase.UpdateMachineData(id, newMachineData)
	if updateError != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(updateError.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (m *machineController) DeleteMachine(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	errDelete := m.usecase.DeleteMachine(id)
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(errDelete.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
