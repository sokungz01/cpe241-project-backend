package usecase

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type maintnanceStatusUsecase struct {
	maintenancestatRepo domain.MaintenanceStatusRepo
}

func NewmaintenanceStatusUsecase(mainmaintenancestatRepo domain.MaintenanceStatusRepo) domain.MaintenanceStatusUsecase {
	return &maintnanceStatusUsecase{maintenancestatRepo: mainmaintenancestatRepo}
}

func (mstatu *maintnanceStatusUsecase) GetAll() (*[]domain.MaintenanceStatus, error) {
	response, err := mstatu.maintenancestatRepo.GetAll()
	return response, err
}
