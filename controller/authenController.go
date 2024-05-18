package controller

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sokungz01/cpe241-project-backend/config"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type AuthController interface {
	SignIn(c *fiber.Ctx) error
	Me(c *fiber.Ctx) error
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
	claims := jwt.MapClaims{
		"userID":     AuthResponse.Id,
		"positionID": AuthResponse.Position,
		"exp":        time.Now().Add(day * 1).Unix(),
	}
	// Create token
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Can't load config", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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

func (u *authenUsecase) Me(c *fiber.Ctx) error {
	reqToken := c.Request().Header.Peek("Authorization")
	splitToken := strings.Split(string(reqToken), "Bearer ")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(splitToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatal("Can't load config", err)
		}
		return []byte(cfg.JWT_ACCESS_TOKEN), nil
	})

	if err != nil {
		return err
	}

	exp, ok := claims["exp"].(float64) // Type assert exp to float64
	if !ok {
		return errors.New("exp claim is not a float64")
	}

	if exp != 0 && int64(exp) < time.Now().Unix() {
		return errors.New("JWT Token is expired")
	}

	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return errors.New("userID claim is not a float64")
	}
	userID := int(userIDFloat)

	AuthResponse, err := u.usecase.Me(int(userID))

	if err != nil {
		return err
	}
	AuthResponse.Password = ""
	return c.JSON(domain.AuthenResponse{
		Data: *AuthResponse,
	})
}
