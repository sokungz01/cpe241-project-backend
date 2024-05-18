package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type UserController interface {
	SignUp(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	Hello(c *fiber.Ctx) error
}

type userUsecase struct {
	usecase domain.UserUseCase
}

func NewUserController(usecase domain.UserUseCase) UserController {
	return &userUsecase{usecase: usecase}
}

func (u *userUsecase) SignUp(c *fiber.Ctx) error {
	newUser := new(domain.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bruh")
	}
	if err := u.usecase.Create(newUser); err != nil {
		return err
	}
	return c.JSON(newUser)
}

func (u *userUsecase) GetAll(c *fiber.Ctx) error {
	users, err := u.usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Pung I here")
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *userUsecase) UpdateUser(c *fiber.Ctx) error {
	id, parseErr := strconv.Atoi(c.Params("id"))
	newUserData := new(domain.User)
	if parseErr != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	parseErrStruct := c.BodyParser(newUserData)
	if parseErrStruct != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseErrStruct.Error())
	}
	response, updateErr := u.usecase.UpdateUser(id, newUserData)
	if updateErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(updateErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *userUsecase) Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).SendString("Hello world")
}
