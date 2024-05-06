package repository

import (
	_"fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type machineRepository struct {
	db *platform.Mysql
}

func NewmachineRepository(db *platform.Mysql) domain.MachineRepository {
	return &machineRepository{db: db}
}

func (m *machineRepository) GetAll() (*[]domain.Machine, error) {
	response := make([]domain.Machine, 0)
	err := m.db.Select(&response, "SELECT *"+
		"FROM `machine`")
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (m *machineRepository) CreateMachineType(mtype domain.MachineType) error {
	_, err := m.db.NamedExec("INSERT INTO `machineType` (`machineTypeName`)"+
		"VALUE (:machineTypeName)", mtype)
	if err != nil {
		return err
	}
	return nil
}

func (m *machineRepository) GetOneMachineTypeByName(typeName string) (*domain.MachineType, error) {
	response := new(domain.MachineType)
	err := m.db.Get(response, "SELECT *"+
		"FROM `machineType`"+
		"WHERE `machineTypeName` = ?", typeName)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineRepository) GetOneMachineTypeByID(id int) (*domain.MachineType, error) {
	response := new(domain.MachineType)
	err := m.db.Get(response, "SELECT *"+
		"FROM `machineType`"+
		"WHERE `machineTypeID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineRepository) UpDateMachineType(id int, newData domain.MachineType) (*domain.MachineType, error) {
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

func (m *machineRepository) DeleteMachineType(id int) error {
	_, err := m.db.Exec("DELETE FROM `machineType` WHERE `machineTypeID` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
