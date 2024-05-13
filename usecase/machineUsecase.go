package usecase

import (
	"errors"
	//"fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
	_ "github.com/sokungz01/cpe241-project-backend/repository"
)

type machineUsecase struct {
	machineRepo        domain.MachineRepository
	machineTypeUsecase domain.MachineTypeUsecase
}

func NewMachineUsecase(machineRepo domain.MachineRepository, machineTypeUsecase domain.MachineTypeUsecase) domain.MachineUsecase {
	return &machineUsecase{machineRepo: machineRepo, machineTypeUsecase: machineTypeUsecase}
}

func (m *machineUsecase) CreateMachine(newMachine *domain.Machine) (*domain.Machine, error) {
	_, err := m.machineTypeUsecase.GetOneMachineTypeByID(newMachine.MachineTypeID)
	if err != nil {
		return nil, errors.New("machine: not valid type or internal server error")
	}
	_, err = m.machineRepo.CreateMachine(newMachine)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *machineUsecase) GetAllMachine() (*[]domain.Machine,error){
	respose,err := m.machineRepo.GetAllMachine()
	if err != nil{
		return nil,err
	}
	return respose,nil
}

func (m *machineUsecase) GetMachineByID(id int) (*domain.Machine,error){
	response,err := m.machineRepo.GetMachineByID(id)
	if err != nil {
		return nil,err
	}
	return response,nil
}