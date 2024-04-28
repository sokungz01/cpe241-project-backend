package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/controller"
	"github.com/sokungz01/cpe241-project-backend/platform"
	"github.com/sokungz01/cpe241-project-backend/repository"
	"github.com/sokungz01/cpe241-project-backend/usecase"
)

func main() {
	//var res *domain.User
	//var err error
	myDB, err := platform.NewSql("root:@tcp(127.0.0.1:3306)/sornrubsom")
	if err != nil {
		fmt.Println(err)
		return
	}
	userRepo := repository.NewUSerRepository(myDB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUseCase)

	api := fiber.New()
	api.Post("/sornrubsom/signup", userController.SignUp)
	api.Listen(":3000")
}
