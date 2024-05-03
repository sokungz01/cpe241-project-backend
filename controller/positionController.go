package controller

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"	
)

type positionController interface{
	CreatePosition(c *fiber.Ctx) error
	GetByPositionName(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type positionUsecase struct{
	posusecase domain.PositionUsecase
}


func NewPositionController(posusecase domain.PositionRepository) positionController{
	return &positionUsecase{posusecase:posusecase}
}

func (usecase *positionUsecase) CreatePosition(c *fiber.Ctx) error{
	newPosition := new(domain.Position)
	if err := c.BodyParser(newPosition);err != nil{
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)

	}
	if err := c.BodyParser(newPosition);err != nil{
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

func (usecase *positionUsecase) GetByPositionName(c *fiber.Ctx)error{
	Parse := new(domain.Position)
	if err := c.BodyParser(Parse); err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	response,err := usecase.posusecase.FindByPositionName(Parse.PositionName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (usecase *positionUsecase) GetAll(c *fiber.Ctx) error{
	response,err := usecase.posusecase.GetAll()
	if err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println(response)
	return c.Status(fiber.StatusOK).JSON(response)
}
