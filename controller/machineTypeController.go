package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type MachineTypeController interface {
	CreateMachineType(c *fiber.Ctx) error
	GetAllMachineType(c *fiber.Ctx) error
	GetOneMachineTypeByID(c *fiber.Ctx) error
	GetOneMachineTypeByName(c *fiber.Ctx) error
	UpdateMachineType(c *fiber.Ctx) error
	DeleteMachineType(c *fiber.Ctx) error
}

type machineTypeUsecase struct {
	usecase domain.MachineTypeUsecase
}

func NewmachineController(usecase domain.MachineTypeUsecase) MachineTypeController {
	return &machineTypeUsecase{usecase: usecase}
}


func (mu *machineTypeUsecase) CreateMachineType(c *fiber.Ctx) error {
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

func (mu *machineTypeUsecase) GetOneMachineTypeByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, err2 := mu.usecase.GetOneMachineTypeByID(id)
	if err2 != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (mu *machineTypeUsecase) GetAllMachineType(c *fiber.Ctx) error {
	response, err := mu.usecase.GetAllMachineType()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (mu *machineTypeUsecase) GetOneMachineTypeByName(c *fiber.Ctx) error {
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

func (mu *machineTypeUsecase) UpdateMachineType(c *fiber.Ctx) error {
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

func (mu *machineTypeUsecase) DeleteMachineType(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := mu.usecase.DeleteMachineType(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
