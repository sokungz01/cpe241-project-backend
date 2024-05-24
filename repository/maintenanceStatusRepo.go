package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type maintenanceStatusRepo struct {
	db *platform.Mysql
}

func NewmaintenanceStatusrepo(db *platform.Mysql) domain.MaintenanceStatusRepo {
	return &maintenanceStatusRepo{db: db}
}

func (mstatr *maintenanceStatusRepo) GetAll() (*[]domain.MaintenanceStatus, error) {
	response := new([]domain.MaintenanceStatus)
	err := mstatr.db.Select(response, "SELECT *"+
		"FROM `maintenanceStatus`")
	if err != nil {
		return nil, err
	}
	return response, nil
}
