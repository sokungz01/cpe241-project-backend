package usecase

import (
	"errors"

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
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*response); i++ {
		(*response)[i].Staff.Password = ""
	}
	return response, nil
}

func (u *maintenanceLogUsecase) GetMaintenanceLogByID(maintainID int) (*domain.MaintenanceLog, error) {
	response, err := u.repo.GetMaintenanceLogByID(maintainID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *maintenanceLogUsecase) GetMaintenanceLogByMachineID(machineID int) (*[]domain.MaintenanceLog, error) {
	response, err := u.repo.GetMaintenanceLogByMachineID(machineID)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*response); i++ {
		(*response)[i].Staff.Password = ""
	}
	return response, nil
}

func (u *maintenanceLogUsecase) GetMaintenanceLogByStaffID(staffID int) (*[]domain.MaintenanceLog, error) {
	response, err := u.repo.GetMaintenanceLogByStaffID(staffID)
	for i := 0; i < len(*response); i++ {

		(*response)[i].Staff.Password = ""
	}
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *maintenanceLogUsecase) CreatemaintenanceLog(newLog *domain.MaintenanceLog) (*domain.MaintenanceLog, error) {
	user, userErr := u.user.GetById(newLog.StaffID)
	if userErr != nil {
		return nil, errors.New("mlogcreate : not a valid staff")
	}
	if user.IsDelete == 1 {
		return nil, errors.New("mlogcreate : not a valid staff")
	}
	_, machineErr := u.machine.GetMachineByID(newLog.MachineID)
	if machineErr != nil {
		return nil, errors.New("mlogcreate : not a valid valid")
	}

	response, err := u.repo.CreatemaintenanceLog(newLog)
	if err != nil {
		return nil, err
	}
	response.Staff.Password = ""
	return response, nil
}

func (u *maintenanceLogUsecase) UpdateMaintenanceLogStatus(maintenanceID int, status int) error {
	if status != 3 && status != 1 {
		return errors.New("value out of range")
	}
	err := u.repo.UpdateMaintenanceLogStatus(maintenanceID, status)
	if err != nil {
		return err
	}
	res, resErr := u.repo.GetMaintenanceLogByID(maintenanceID)
	if resErr != nil {
		return resErr
	}
	mStatErr := u.machine.UpdateMachineStatus(res.MachineID, 3)
	if mStatErr != nil {
		return mStatErr
	}
	return nil
}
