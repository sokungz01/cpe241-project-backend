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

	itemRepo := repository.NewItemRepository(myDB)

	itemLogRepo := repository.NewItemLogRepository(myDB)
	itemLogUsecase := usecase.NewItemLogUsecase(itemLogRepo, itemRepo, userRepo)
	itemLogController := controller.NewItemLogController(itemLogUsecase)

	itemUsecase := usecase.NewItemUsecase(itemRepo, itemLogUsecase)
	itemController := controller.NewItemController(itemUsecase)

	errorLogRepo := repository.NewErrorLogRepository(myDB)
	errorTypeRepo := repository.NewErrorTypeRepository(myDB)
	errorTypeUsecase := usecase.NewErrorTypeUsecase(errorTypeRepo)
	errorTypeController := controller.NewErrorTypeController(errorTypeUsecase)

	maintenancePartsRepo := repository.NewMaintenancePartsRepository(myDB)

	serviceRequestRepo := repository.NewServiceRequestRepository(myDB)
	serviceRequestUsecase := usecase.NewServiceRequestUsecase(serviceRequestRepo, userUseCase, machineUsecase, errorTypeUsecase, errorLogRepo)
	serviceRequestController := controller.NewServiceRequestController(serviceRequestUsecase)

	serviceResponseRepo := repository.NewServiceResponseRepository(myDB)
	serviceResponseUsecase := usecase.NewServiceResponsUsecase(serviceResponseRepo, userUseCase, serviceRequestUsecase, itemUsecase, itemRepo, maintenancePartsRepo, itemLogUsecase, machineUsecase)
	serviceResponseController := controller.NewServiceResponseController(serviceResponseUsecase)

	maintenanceStatusRepo := repository.NewmaintenanceStatusrepo(myDB)
	maintenanceStatusUsecase := usecase.NewmaintenanceStatusUsecase(maintenanceStatusRepo)
	maintenanceStatusController := controller.NewMaintenanceStatuscontroller(maintenanceStatusUsecase)

	maintenanceLogRepo := repository.NewMaintenanceLogRepository(myDB)
	maintenanceLogUsecase := usecase.NewMaintenanceLogUsecase(maintenanceLogRepo, userUseCase, machineUsecase)
	maintenanceLogController := controller.NewMaintenanceLogController(maintenanceLogUsecase)

	authGroup := app.Group("/auth")
	authGroup.Get("/me", jwt, authController.Me)
	authGroup.Post("/signup", userController.SignUp)
	authGroup.Post("/signin", authController.SignIn)
	authGroup.Put("/update/:id", jwt, userController.UpdateUser)

	userGroup := app.Group("/user")
	userGroup.Get("/all", jwt, userController.GetAll)
	userGroup.Get("/:id", jwt, userController.GetByUserID)
	userGroup.Delete("/delete/:id", jwt, userController.DeleteUser)

	positionGroup := app.Group("/position")
	positionGroup.Get("/", jwt, positionController.GetAll)
	positionGroup.Get("/findbypositionname", jwt, positionController.GetByPositionName)
	positionGroup.Get("/:id", jwt, positionController.GetByPositionID)
	positionGroup.Put("/:id", jwt, positionController.UpdatePosition)
	positionGroup.Post("/", jwt, positionController.CreatePosition)
	positionGroup.Delete("/:id", jwt, positionController.DeletePosition)

	machineGroup := app.Group("/machine")
	machineGroup.Get("/", machineController.GetAllMachine)
	machineGroup.Get("/getbyname", machineController.GetMachineByName)
	machineGroup.Get("/:id", machineController.GetMachineByID)
	machineGroup.Put("/updatestatus/:id", machineController.UpdateMachineStatus)
	machineGroup.Put("/:id", machineController.UpdateMachineData)
	machineGroup.Post("/", machineController.CreateMachine)
	machineGroup.Delete("/:id", machineController.DeleteMachine)

	machineTypeGroup := app.Group("/machinetype")
	machineTypeGroup.Get("/", jwt, machineTypeController.GetAllMachineType)
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

	item := app.Group("/item")
	item.Get("/", itemController.GetAllItem)
	item.Get("/:id", itemController.FindByID)
	item.Post("/", itemController.CreateItem)
	item.Put("/:id", itemController.UpdateItem)
	item.Delete("/:id", itemController.DeleteItem)

	itemLog := app.Group("/itemLog")
	itemLog.Get("/", itemLogController.GetAll)
	itemLog.Post("/", itemLogController.CreateItemLog)

	errorType := app.Group("/errortype")
	errorType.Get("/", errorTypeController.GetAllErrorType)
	errorType.Get("/:id", errorTypeController.FindByID)
	errorType.Post("/", errorTypeController.CreateErrorType)
	errorType.Put("/:id", errorTypeController.UpdateErrorType)
	errorType.Delete("/:id", errorTypeController.DeleteErrorType)

	serviceRequest := app.Group("/servicerequest")
	serviceRequest.Get("/", serviceRequestController.GetAllServiceRequest)
	serviceRequest.Get("/:id", serviceRequestController.GetServiceRequest)
	serviceRequest.Post("/", serviceRequestController.CreateServiceRequest)
	serviceRequest.Put("/status/:id", serviceRequestController.UpdateServiceRequestStatus)

	serviceResponse := app.Group("/serviceresponse")
	serviceResponse.Get("/", serviceResponseController.GetAll)
	serviceResponse.Get("/:id", serviceResponseController.GetOne)
	serviceResponse.Get("/service/:id", serviceResponseController.GetOneByService)
	serviceResponse.Post("/", serviceResponseController.CreateResponse)
	serviceResponse.Delete("/:id", serviceResponseController.DeleteResponse)

	maintenanceStat := app.Group("/maintenancestatus")
	maintenanceStat.Get("/", maintenanceStatusController.GetAll)

	maintenanceLog := app.Group("/maintenancelog")
	maintenanceLog.Get("/", maintenanceLogController.GetAllmaintenanceLog)
	maintenanceLog.Get("/getbymachineID/:id", maintenanceLogController.GetMaintenanceLogByMachineID)
	maintenanceLog.Get("/getbystaffID/:id", maintenanceLogController.GetMaintenanceLogByStaffID)
	maintenanceLog.Post("/", maintenanceLogController.CreatemaintenanceLog)
	maintenanceLog.Put("/:id", maintenanceLogController.UpdateMaintenanceLogStatus)

	analysisRepo := repository.NewAnalysisRepo(myDB)
	analysisController := controller.NewAnalysisController(analysisRepo)

	analysis := app.Group("/analyze")
	analysis.Get("/timeseriesInventory", analysisController.GetInventoryAnanlysis)
	analysis.Get("/machinetypeerror", analysisController.GetMachineTypeErrorAnalysis)
	analysis.Get("/employeeengage", analysisController.GetEmployeeEngagementAnalysis)
	analysis.Get("/maintenancecos", analysisController.GetMaintenanceCostAnalysis)
}
