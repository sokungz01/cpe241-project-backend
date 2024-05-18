package repository

import (
	_ "fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type machineRepository struct {
	db *platform.Mysql
}

func NewmachineRepository(db *platform.Mysql) domain.MachineRepository {
	return &machineRepository{db: db}
}

func (m *machineRepository) CreateMachine(newMachine *domain.Machine) (*domain.Machine, error) {
	_, err := m.db.NamedExec("INSERT INTO `machine` (`machineName`,`machineBrand`,`machineTypeID`,`startDate`,`description`,`status`)"+
		"VALUE (:machineName,:machineBrand,:machineTypeID,:startDate,:description,:status)", newMachine)
	if err != nil {
		return nil, err
	}
	return newMachine, nil
}

func (m *machineRepository) GetAllMachine() (*[]domain.Machine, error) {

	response := make([]domain.Machine, 0)
	if err := m.db.Select(&response, "SELECT * FROM `machine`"); err != nil {
		return nil, err
	}
	return &response, nil
}

func (m *machineRepository) GetMachineByID(id int) (*domain.Machine, error) {
	response := new(domain.Machine)
	err := m.db.Get(response, "SELECT *"+
		"FROM `machine`"+
		"WHERE `machineID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineRepository) GetMachineByName(machineName string) (*[]domain.Machine, error) {
	response := new([]domain.Machine)
	queryStr := "'%" + machineName + "%'"
	err := m.db.Select(response, "SELECT *"+
		"FROM `machine`"+
		"WHERE `machineName` LIKE "+queryStr)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *machineRepository) UpdateMachineData(id int, newMachineData *domain.Machine) error {
	_, err := m.db.Exec("UPDATE `machine`"+
		"SET `machineName`= ?,`machineBrand`= ?,`machineTypeID`= ?,`startDate`= ?,`endDate`= ?,`description`= ?,`status`= ? WHERE `machineID`= ?", newMachineData.MachineName, newMachineData.MachineBrand, newMachineData.MachineTypeID, newMachineData.StartDate, newMachineData.EndDate, newMachineData.Description, newMachineData.Status, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *machineRepository) DeleteMachine(id int) error {
	_, err := m.db.Exec("DELETE FROM `machine WHERE `machineID` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
