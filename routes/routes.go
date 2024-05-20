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

	machineTypeRepo := repository.NewmachineTypeRepository(myDB)
	machineTypeUsecase := usecase.NewMachineTypeUsecase(machineTypeRepo)
	machineTypeController := controller.NewmachineController(machineTypeUsecase)

	machineRepo := repository.NewmachineRepository(myDB)
	machineUsecase := usecase.NewMachineUsecase(machineRepo, machineTypeUsecase)
	machineController := controller.NewMachineController(machineUsecase)

	itemCategoryRepo := repository.NewItemCategoryRepository(myDB)
	itemCategoryUsecase := usecase.NewItemCategoryUsecase(itemCategoryRepo)
	itemCategoryController := controller.NewItemCategoryController(itemCategoryUsecase)

	authGroup := app.Group("/auth")
	authGroup.Get("/me", jwt, authController.Me)
	authGroup.Post("/signup", userController.SignUp)
	authGroup.Post("/signin", authController.SignIn)
	authGroup.Put("/update/:id", jwt, userController.UpdateUser)

	userGroup := app.Group("/user")
	userGroup.Get("/all", jwt, userController.GetAll)
	userGroup.Get("/:id", jwt, userController.GetByUserID)

	positionGroup := app.Group("/position")
	positionGroup.Get("/", positionController.GetAll)
	positionGroup.Get("/findbypositionname", positionController.GetByPositionName)
	positionGroup.Post("/", positionController.CreatePosition)

	machineGroup := app.Group("/machine")
	machineGroup.Get("/", machineController.GetAllMachine)
	machineGroup.Get("/getbyname", machineController.GetMachineByName)
	machineGroup.Get("/:id", machineController.GetMachineByID)
	machineGroup.Put("/:id", machineController.UpdateMachineData)
	machineGroup.Post("/", machineController.CreateMachine)
	machineGroup.Delete("/:id", machineController.DeleteMachine)

	machineTypeGroup := app.Group("/machinetype")
	machineTypeGroup.Get("/", machineTypeController.GetAllMachineType)
	machineTypeGroup.Get("/:id", machineTypeController.GetOneMachineTypeByID)
	machineTypeGroup.Get("/getbyname", machineTypeController.GetOneMachineTypeByName)
	machineTypeGroup.Post("/", machineTypeController.CreateMachineType)
	machineTypeGroup.Put("/:id", machineTypeController.UpdateMachineType)
	machineTypeGroup.Delete("/:id", machineTypeController.DeleteMachineType)

	itemCategory := app.Group("/itemCategory")
	itemCategory.Get("/", itemCategoryController.GetAllItemCategory)
	itemCategory.Get("/:id", itemCategoryController.FindByID)
	itemCategory.Post("/", itemCategoryController.CreateItemCategory)
	itemCategory.Put("/:id", itemCategoryController.UpdateItemCategory)
	itemCategory.Delete("/:id", itemCategoryController.DeleteItemCategory)

}
