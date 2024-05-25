package domain

import (
	"time"
)

type MaintenanceLog struct {
	MaintenanceID int       `json:"maintenanceID" db:"maintenanceID"`
	StaffID       int       `json:"staffID" db:"staffID"`
	Staff         User      `json:"staff" db:",prefix=employee."`
	MachineID     int       `json:"machineID" db:"machineID"`
	Period        int       `json:"period" db:"period"`
	CreatedDate   time.Time `json:"createdDate" db:"createdDate"`
	UpdateDate    time.Time `json:"updateDate" db:"updateDate"`
	StatusID      int       `json:"statusID" db:"statusID"`
}

type MaintenanceLogRepository interface {
	GetAllmaintenanceLog() (*[]MaintenanceLog, error)
}

type MaintenanceLogUsecase interface {
	GetAllmaintenanceLog() (*[]MaintenanceLog, error)
}
