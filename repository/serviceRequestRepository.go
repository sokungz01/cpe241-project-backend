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
		"FROM `serviceRequest`")
	if err != nil {
		return nil, err
	}
	return &response, nil
}
