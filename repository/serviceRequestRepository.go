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

func (r *serviceRequestRepository) CreateServiceRequest(newServiceRequest *domain.ServiceRequest) (*domain.ServiceRequest, error) {
	_, err := r.db.NamedExec("INSERT INTO `serviceRequest`(`employeeID`, `machineID`, `description`, `createdDate`, `updateDate`, `statusID`) "+
		"VALUES (:employeeID, :machineID, :description, :createdDate, :createdDate, 1)", newServiceRequest)
	if err != nil {
		return nil, err
	}

	response := new(domain.ServiceRequest)
	_ = r.db.Get(response, "SELECT * FROM `serviceRequest` WHERE serviceID IN (SELECT LAST_INSERT_ID() as id)")

	return response, nil
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

func (r *serviceRequestRepository) UpdateServiceRequestStatus(id int, statusID int) (*domain.ServiceRequest, error) {
	_, err := r.db.Exec("UPDATE `serviceRequest`"+
		"SET `statusID`= ? WHERE `serviceID`= ? ", statusID, id)
	if err != nil {
		return nil, err
	}
	response, _ := r.GetServiceRequest(id)
	return response, nil
}
