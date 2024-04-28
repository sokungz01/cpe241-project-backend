package controller

import (
	//_"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
	//"github.com/sokungz01/cpe241-project-backend/usecase"
)
type UserController interface{
	SignUp(c *fiber.Ctx) error
}

type userUsecase struct {
	usecase domain.UserUseCase
}

func NewUserController(usecase domain.UserUseCase) UserController{
	return &userUsecase{usecase : usecase}
}

func (u *userUsecase)SignUp(c *fiber.Ctx) error {
	newUser := new(domain.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bruh")
	}
	newUser.Password = "HumYaiYai"
	if err := u.usecase.Create(newUser); err != nil{
		return err
	}
	return c.JSON(newUser)
}
