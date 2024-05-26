package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type maintenanceLogrepository struct {
	db *platform.Mysql
}

func NewMaintenanceLogRepository(db *platform.Mysql) domain.MaintenanceLogRepository {
	return &maintenanceLogrepository{db: db}
}

func (r *maintenanceLogrepository) GetAllmaintenanceLog() (*[]domain.MaintenanceLog, error) {
	response := new([]domain.MaintenanceLog)
	err := r.db.Select(response, "SELECT * "+
		"FROM `maintenanceLog` "+
		"INNER JOIN employee ON maintenanceLog.staffID = employee.employeeID "+
		"INNER JOIN machine ON maintenanceLog.machineID = machine.machineID")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *maintenanceLogrepository) GetMaintenanceLogByID(maintainID int) (*domain.MaintenanceLog, error) {
	response := new(domain.MaintenanceLog)
	err := r.db.Get(response, "SELECT * "+
		"FROM `maintenanceLog` "+
		"INNER JOIN employee ON maintenanceLog.staffID = employee.employeeID "+
		"INNER JOIN machine ON maintenanceLog.machineID = machine.machineID "+
		"WHERE `maintenanceLog`.`maintenanceID` = ?", maintainID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *maintenanceLogrepository) GetMaintenanceLogByMachineID(machineID int) (*[]domain.MaintenanceLog, error) {
	response := new([]domain.MaintenanceLog)
	err := r.db.Select(response, "SELECT * "+
		"FROM `maintenanceLog` "+
		"INNER JOIN employee ON maintenanceLog.staffID = employee.employeeID "+
		"INNER JOIN machine ON maintenanceLog.machineID = machine.machineID "+
		"WHERE `maintenanceLog`.`machineID` = ?", machineID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *maintenanceLogrepository) GetMaintenanceLogByStaffID(staffID int) (*[]domain.MaintenanceLog, error) {
	response := new([]domain.MaintenanceLog)
	err := r.db.Select(response, "SELECT * "+
		"FROM `maintenanceLog` "+
		"INNER JOIN employee ON maintenanceLog.staffID = employee.employeeID "+
		"INNER JOIN machine ON maintenanceLog.machineID = machine.machineID "+
		"WHERE `maintenanceLog`.`staffID` = ?", staffID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *maintenanceLogrepository) CreatemaintenanceLog(newLog *domain.MaintenanceLog) (*domain.MaintenanceLog, error) {
	response := new(domain.MaintenanceLog)
	_, err := r.db.Exec("INSERT INTO `maintenanceLog` (`staffID`,`machineID`,`period`,`createdDate`,`updateDate`,`statusID`) "+
		"VALUE (?,?,?,?,?,?)",
		newLog.StaffID, newLog.MachineID, newLog.Period, newLog.CreatedDate, newLog.UpdateDate, newLog.StatusID)
	if err != nil {
		return nil, err
	}
	err = r.db.Get(response, "SELECT * "+
		"FROM `maintenanceLog` "+
		"INNER JOIN employee ON maintenanceLog.staffID = employee.employeeID "+
		"INNER JOIN machine ON maintenanceLog.machineID = machine.machineID "+
		"WHERE `maintenanceID` IN  (SELECT LAST_INSERT_ID() as id)")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *maintenanceLogrepository) UpdateMaintenanceLogStatus(maintenanceID int, status int) error {

	_, err := r.db.Exec("UPDATE `maintenanceLog` "+
		"SET `statusID` = ? "+
		"WHERE `maintenanceID` = ?", status, maintenanceID)
	return err
}
