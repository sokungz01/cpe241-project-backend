package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type machineTypeUsecase struct {
	repo domain.MachineTypeRepository
}

func NewMachineTypeUsecase(repo domain.MachineTypeRepository) domain.MachineTypeUsecase {
	return &machineTypeUsecase{repo: repo}
}


func (mr *machineTypeUsecase) CreateMachineType(mtype domain.MachineType) error {

	resp, _ := mr.GetOneMachineTypeByName(mtype.MachineTypeName)
	if resp != nil {
		return errors.New("machinetype : already existed type")
	}
	if err := mr.repo.CreateMachineType(mtype); err != nil {
		return err
	}
	return nil
}

func (mr *machineTypeUsecase) GetAllMachineType() (*[]domain.MachineType,error){
	response,err := mr.repo.GetAllMachineType()
	if err != nil{
		return nil,err
	}
	return response,nil
}

func (mr *machineTypeUsecase) GetOneMachineTypeByName(typeName string) (*domain.MachineType, error) {
	response, err := mr.repo.GetOneMachineTypeByName(typeName)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (mr *machineTypeUsecase) GetOneMachineTypeByID(id int) (*domain.MachineType, error){
	response,err := mr.repo.GetOneMachineTypeByID(id)
	if err != nil{
		return nil,err
	}
	return response,nil
}

func (mr *machineTypeUsecase) UpDateMachineType(id int, newData domain.MachineType) (*domain.MachineType, error) {
	response, err := mr.repo.UpDateMachineType(id, newData)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (mr *machineTypeUsecase) DeleteMachineType(id int) error {

	if _, err := mr.repo.GetOneMachineTypeByID(id); err != nil {
		return err
	}
	err := mr.repo.DeleteMachineType(id)
	if err != nil {
		return err
	}
	return nil
}
