package usecase

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type maintenanceLogUsecase struct {
	repo    domain.MaintenanceLogRepository
	user    domain.UserUseCase
	machine domain.MachineUsecase
}

func NewMaintenanceLogUsecase(repo domain.MaintenanceLogRepository,
	user domain.UserUseCase, machine domain.MachineUsecase) domain.MaintenanceLogUsecase {
	return &maintenanceLogUsecase{
		repo:    repo,
		user:    user,
		machine: machine,
	}
}

func (u *maintenanceLogUsecase) GetAllmaintenanceLog() (*[]domain.MaintenanceLog, error) {
	response, err := u.repo.GetAllmaintenanceLog()
	for i, _ := range *response {
		(*response)[i].Staff.Password = ""
	}
	if err != nil {
		return nil, err
	}
	return response, nil
}
