package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type serviceRequestRepository struct {
	db *platform.Mysql
}

func NewServiceRequestRepository(db *platform.Mysql) domain.ServiceRequestRepository {
	return &serviceRequestRepository{db: db}
}

func (r *serviceRequestRepository) GetAllServiceRequest() (*[]domain.ServiceRequest, error) {
	response := make([]domain.ServiceRequest, 0)
	err := r.db.Select(&response, "SELECT *"+
		"FROM `serviceRequest`"+
		"INNER JOIN employee ON employee.employeeID = serviceRequest.`employeeID`"+
		"INNER JOIN machine ON machine.machineID = serviceRequest.`machineID`")
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *serviceRequestRepository) GetServiceRequest(id int) (*domain.ServiceRequest, error) {
	response := new(domain.ServiceRequest)
	err := r.db.Get(response, "SELECT *"+
		"FROM `serviceRequest`"+
		"INNER JOIN employee ON employee.employeeID = serviceRequest.`employeeID`"+
		"INNER JOIN machine ON machine.machineID = serviceRequest.`machineID` WHERE `serviceID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
