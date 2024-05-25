package domain

import (
	"database/sql"
	"time"
)

type ServiceRequest struct {
	ServiceID   int          `json:"serviceID" db:"serviceID"`
	EmployeeID  int          `json:"employeeID" db:"employeeID"`
	Employee    User         `json:"user" db:",prefix=employee."`
	MachineID   int          `json:"machineID" db:"machineID"`
	Machine     Machine      `json:"machine" db:",prefix=machine."`
	ErrorLog    []ErrorLog   `json:"errorLog" db:"errorLog"`
	Description string       `json:"description" db:"description"`
	CreatedDate time.Time    `json:"createdDate" db:"createdDate"`
	UpdateDate  sql.NullTime `json:"updateDate" db:"updateDate"`
	StatusID    int          `json:"statusID" db:"statusID"`
}

type ServiceRequestRepository interface {
	GetAllServiceRequest() (*[]ServiceRequest, error)
	GetServiceRequest(id int) (*ServiceRequest, error)
}

type ServiceRequestUsecase interface {
	GetAllServiceRequest() (*[]ServiceRequest, error)
	GetServiceRequest(id int) (*ServiceRequest, error)
	CreateServiceRequest(newServiceRequest *ServiceRequest) error
}
