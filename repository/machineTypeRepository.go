package repository

import (
	_"fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type machineTypeRepository struct {
	db *platform.Mysql
}

func NewmachineTypeRepository(db *platform.Mysql) domain.MachineTypeRepository {
	return &machineTypeRepository{db: db}
}


func (m *machineTypeRepository) CreateMachineType(mtype domain.MachineType) error {
	_, err := m.db.NamedExec("INSERT INTO `machineType` (`machineTypeName`)"+
		"VALUE (:machineTypeName)", mtype)
	if err != nil {
		return err
	}
	return nil
}

func (m *machineTypeRepository) GetAllMachineType() (*[]domain.MachineType,error) {
	response := make([]domain.MachineType,0)
	if err := m.db.Select(&response,"SELECT * FROM `machineType`");err != nil{
		return nil,err
	}
	return &response,nil
}

func (m *machineTypeRepository) GetOneMachineTypeByName(typeName string) (*domain.MachineType, error) {
	response := new(domain.MachineType)
	err := m.db.Get(response, "SELECT *"+
		"FROM `machineType`"+
		"WHERE `machineTypeName` = ?", typeName)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineTypeRepository) GetOneMachineTypeByID(id int) (*domain.MachineType, error) {
	response := new(domain.MachineType)
	err := m.db.Get(response, "SELECT *"+
		"FROM `machineType`"+
		"WHERE `machineTypeID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineTypeRepository) UpDateMachineType(id int, newData domain.MachineType) (*domain.MachineType, error) {
	response := new(domain.MachineType)
	_, err := m.db.Exec("UPDATE `machineType`"+
		"SET `machineTypeName` = ? WHERE `machineTypeID` = ?", newData.MachineTypeName, id)
	if err != nil {
		return nil, err
	}
	//query updated data for response
	err = m.db.Get(response, "SELECT * FROM `machineType` WHERE `machineTypeID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineTypeRepository) DeleteMachineType(id int) error {
	_, err := m.db.Exec("DELETE FROM `machineType` WHERE `machineTypeID` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
