package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/sokungz01/cpe241-project-backend/config"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type AuthController interface {
	SignIn(c *fiber.Ctx) error
}

type authenUsecase struct {
	usecase domain.AuthenUsecase
}

func NewAuthenController(usecase domain.AuthenUsecase) AuthController {
	return &authenUsecase{usecase: usecase}
}

func (u *authenUsecase) SignIn(c *fiber.Ctx) error {
	newUser := new(domain.AuthenPayload)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Missing body")
	}
	AuthResponse, err := u.usecase.SignIn(newUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"userID":     AuthResponse.ID,
		"positionID": AuthResponse.Position,
		"exp":        time.Now().Add(day * 1).Unix(),
	}
	// Create token
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Can't load config", err)
	}
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(cfg.JWT_ACCESS_TOKEN))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(domain.AuthenResponse{
		Token: t,
	})
}
