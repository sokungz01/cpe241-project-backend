package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type MachineController interface {
	GetAllMachine(c *fiber.Ctx) error

	CreateMachineType(c *fiber.Ctx) error
	GetOneMachineTypeByName(c *fiber.Ctx) error
	UpdateMachineType(c *fiber.Ctx) error
	DeleteMachineType(c *fiber.Ctx) error
}

type machineUseCase struct {
	usecase domain.MachineUseCase
}

func NewmachineController(usecase domain.MachineUseCase) MachineController {
	return &machineUseCase{usecase: usecase}
}

func (mu *machineUseCase) GetAllMachine(c *fiber.Ctx) error {
	return c.SendString("Kuay")
}

func (mu *machineUseCase) CreateMachineType(c *fiber.Ctx) error {
	newMachineType := new(domain.MachineType)
	var err error
	if err := c.BodyParser(newMachineType); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := mu.usecase.CreateMachineType(*newMachineType); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	newMachineType, err = mu.usecase.GetOneMachineTypeByName(newMachineType.MachineTypeName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(newMachineType)
}

func (mu *machineUseCase) GetOneMachineTypeByName(c *fiber.Ctx) error {
	parse := new(domain.MachineType)
	if err := c.BodyParser(parse); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err := mu.usecase.GetOneMachineTypeByName(parse.MachineTypeName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (mu *machineUseCase) UpdateMachineType(c *fiber.Ctx) error {
	Data := new(domain.MachineType)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := c.BodyParser(Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	Data, err = mu.usecase.UpDateMachineType(id, *Data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(Data)
}

func (mu *machineUseCase) DeleteMachineType(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := mu.usecase.DeleteMachineType(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
