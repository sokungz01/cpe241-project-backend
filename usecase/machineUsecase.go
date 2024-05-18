package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
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

func (m *machineUsecase) GetAllMachine() (*[]domain.Machine, error) {
	respose, err := m.machineRepo.GetAllMachine()
	if err != nil {
		return nil, err
	}
	return respose, nil
}

func (m *machineUsecase) GetMachineByID(id int) (*domain.Machine, error) {
	response, err := m.machineRepo.GetMachineByID(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineUsecase) DeleteMachine(id int) error {
	_, err := m.GetMachineByID(id)
	if err != nil {
		return err
	}
	return m.machineRepo.DeleteMachine(id)
}

func (m *machineUsecase) GetMachineByName(machineName string) (*domain.Machine, error) {
	response, err := m.machineRepo.GetMachineByName(machineName)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineUsecase) UpdateMachineData(id int, newMachineData *domain.Machine) (*domain.Machine, error) {
	_, getErr := m.GetMachineByID(id)
	if getErr != nil {
		return nil, getErr
	}
	updateErr := m.machineRepo.UpdateMachineData(id, newMachineData)
	if updateErr != nil {
		return nil, updateErr
	}
	response, checkoutGetErr := m.GetMachineByID(id)
	if checkoutGetErr != nil {
		return nil, checkoutGetErr
	}
	return response, nil
}
