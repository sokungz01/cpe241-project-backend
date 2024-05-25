package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type serviceResponseReponseitory struct {
	db *platform.Mysql
}

func NewServiceResponseRepository(db *platform.Mysql) domain.ServiceResponseRepository {
	return &serviceResponseReponseitory{db: db}
}

func (r *serviceResponseReponseitory) GetAllResponse() (*[]domain.ServiceResponse, error) {
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
