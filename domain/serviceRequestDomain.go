package domain

import (
	"database/sql"
	"time"
)

type ServiceRequest struct {
	ServiceID   int          `json:"serviceID" db:"serviceID"`
	EmployeeID  int          `json:"employeeID" db:"employeeID"`
	MachineID   int          `json:"machineID" db:"machineID"`
	ErrorLog    []ErrorLog   `json:"errorLog" db:"errorLog"`
	Description string       `json:"description" db:"description"`
	CreatedDate time.Time    `json:"createdDate" db:"createdDate"`
	UpdateDate  sql.NullTime `json:"updateDate" db:"updateDate"`
}

type ServiceRequestRepository interface {
	GetAllServiceRequest() (*[]ServiceRequest, error)
}

type ServiceRequestUsecase interface {
	GetAllServiceRequest() (*[]ServiceRequest, error)
	CreateServiceRequest(newServiceRequest *ServiceRequest) error
}
