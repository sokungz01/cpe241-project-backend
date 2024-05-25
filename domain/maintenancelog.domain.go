package domain

import (
	"time"
)

type MaintenanceLog struct {
	MaintenanceID int       `json:"maintenanceID" db:"maintenanceID"`
	StaffID       int       `json:"staffID" db:"staffID"`
	Staff         User      `json:"staff" db:",prefix=employee."`
	MachineID     int       `json:"machineID" db:"machineID"`
	Machine       Machine   `json:"machine" db:",prefix=machine."`
	Period        int       `json:"period" db:"period"`
	CreatedDate   time.Time `json:"createdDate" db:"createdDate"`
	UpdateDate    time.Time `json:"updateDate" db:"updateDate"`
	StatusID      int       `json:"statusID" db:"statusID"`
}

type MaintenanceLogRepository interface {
	GetAllmaintenanceLog() (*[]MaintenanceLog, error)
	GetMaintenanceLogByMachineID(machineID int) (*[]MaintenanceLog, error)
	GetMaintenanceLogByStaffID(staffID int) (*[]MaintenanceLog, error)
	CreatemaintenanceLog(newLog *MaintenanceLog) (*MaintenanceLog, error)
	UpdateMaintenanceLogStatus(maintenanceID int, status int) error
}

type MaintenanceLogUsecase interface {
	GetAllmaintenanceLog() (*[]MaintenanceLog, error)
	GetMaintenanceLogByMachineID(machineID int) (*[]MaintenanceLog, error)
	GetMaintenanceLogByStaffID(staffID int) (*[]MaintenanceLog, error)
	CreatemaintenanceLog(newLog *MaintenanceLog) (*MaintenanceLog, error)
	UpdateMaintenanceLogStatus(maintenanceID int, status int) error
}
