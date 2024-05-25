package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type serviceResponseRepository struct {
	db *platform.Mysql
}

func NewServiceResponseRepository(db *platform.Mysql) domain.ServiceResponseRepository {
	return &serviceResponseRepository{db: db}
}

func (r *serviceResponseRepository) GetAllResponse() (*[]domain.ServiceResponse, error) {
	response := new([]domain.ServiceResponse)
	err := r.db.Select(response, "SELECT *,`serviceResponse`.`description` AS `desc`"+
		"FROM `serviceResponse`"+
		"INNER JOIN employee ON employee.employeeID = `serviceResponse`.`staffID`"+
		"INNER JOIN serviceRequest ON serviceRequest.serviceID = serviceResponse.requestedServiceID")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *serviceResponseRepository) CreateServiceResponse(newResponse *domain.ServiceResponse) error {
	_, err := r.db.Exec("INSERT INTO `serviceResponse` (`staffID`,`requestedServiceID`,`title`,`description`,`createdDate`,`updateDate`)"+
		"VALUES (?,?,?,?,?,?)", newResponse.StaffID, newResponse.RequestedServiceID, newResponse.Title, newResponse.Description, newResponse.CreatedDate, newResponse.UpdateDate)
	return err
}

func (r *serviceResponseRepository) DeleteResponse(id int) error {
	_, err := r.db.Exec("DELETE FROM `serviceResponse`"+
		" WHERE `serviceResponse`.`staffServiceID` = ?", id)
	return err
}

func (r *serviceResponseRepository) GetResponse(id int) (*domain.ServiceResponse, error) {
	response := new(domain.ServiceResponse)
	err := r.db.Get(response, "SELECT *,`serviceResponse`.`description` AS `desc`"+
		"FROM `serviceResponse`"+
		"INNER JOIN employee ON employee.employeeID = `serviceResponse`.`staffID`"+
		"INNER JOIN serviceRequest ON serviceRequest.serviceID = serviceResponse.requestedServiceID"+
		" WHERE `serviceResponse`.`staffServiceID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *serviceResponseRepository) GetResponseByService(id int) (*[]domain.ServiceResponse, error) {
	response := make([]domain.ServiceResponse, 0)
	err := r.db.Select(&response, "SELECT *,`serviceResponse`.`description` AS `desc`"+
		"FROM `serviceResponse`"+
		"INNER JOIN employee ON employee.employeeID = `serviceResponse`.`staffID`"+
		"INNER JOIN serviceRequest ON serviceRequest.serviceID = serviceResponse.requestedServiceID"+
		" WHERE `serviceResponse`.`requestedServiceID` = ?", id)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
