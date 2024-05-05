package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/config"
	"github.com/sokungz01/cpe241-project-backend/controller"
	"github.com/sokungz01/cpe241-project-backend/middleware"
	"github.com/sokungz01/cpe241-project-backend/platform"
	"github.com/sokungz01/cpe241-project-backend/repository"
	"github.com/sokungz01/cpe241-project-backend/usecase"
)

func RoutesRegister(app *fiber.App, myDB *platform.Mysql, cfg *config.Config) {
	jwt := middleware.NewAuthMiddleware(cfg.JWT_ACCESS_TOKEN)
	userRepo := repository.NewUSerRepository(myDB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUseCase)

	authRepo := repository.NewAuthenRepository(myDB)
	authUseCase := usecase.NewAuthUseCase(authRepo)
	authController := controller.NewAuthenController(authUseCase)

	positionRepo := repository.NewPositionRepository(myDB)
	positionUsecase := usecase.NewPositionUsecase(positionRepo)
	positionController := controller.NewPositionController(positionUsecase)
	
	authGroup := app.Group("/auth")
	authGroup.Get("/me", jwt, authController.Me)
	authGroup.Post("/signup", userController.SignUp)
	authGroup.Post("/signin", authController.SignIn)

	userGroup := app.Group("/user")
	userGroup.Get("/all", jwt, userController.GetAll)

	positionGroup := app.Group("/position")
	positionGroup.Get("/",positionController.GetAll)
	positionGroup.Get("/findbypositionname",positionController.GetByPositionName)
	positionGroup.Post("/",positionController.CreatePosition)
}
