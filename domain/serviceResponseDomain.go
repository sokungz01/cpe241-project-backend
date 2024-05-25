package domain

import (
	"time"
)

type ServiceResponse struct {
	StaffServiceID     int            `json:"staffServiceID" db:"staffServiceID"`
	StaffID            int            `json:"staffID" db:"staffID"`
	Staff              User           `json:"user" db:",prefix=employee."`
	RequestedServiceID int            `json:"requestedServiceID" db:"requestedServiceID"`
	RequestedService   ServiceRequest `json:"serviceRequest" db:",prefix=serviceRequest."`
	Title              string         `json:"title" db:"title"`
	Description        string         `json:"description" db:"desc"`
	CreatedDate        time.Time      `json:"createdDate" db:"createdDate"`
	UpdateDate         time.Time      `json:"updateDate" db:"updateDate"`
}

type ServiceResponseRepository interface {
	GetAllResponse() (*[]ServiceResponse, error)
	GetResponse(id int) (*ServiceResponse, error)
	GetResponseByService(id int) (*[]ServiceResponse, error)
	DeleteResponse(id int) error
	CreateServiceResponse(newResponse *ServiceResponse) error
}

type ServiceResponseUsecase interface {
	GetAllResponse() (*[]ServiceResponse, error)
	GetResponse(id int) (*ServiceResponse, error)
	GetResponseByService(id int) (*[]ServiceResponse, error)
	DeleteResponse(id int) error
	CreateServiceResponse(newResponse *ServiceResponse) error
}
