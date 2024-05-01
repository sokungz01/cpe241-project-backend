package controller

import (
	//_"fmt"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
	//"github.com/sokungz01/cpe241-project-backend/usecase"
)

type UserController interface {
	SignUp(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
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
	users ,err := u.usecase.GetAll()
	fmt.Println(err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Pung I here")
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *userUsecase) Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).SendString("Hello world")
}
