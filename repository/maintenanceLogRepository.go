package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type maintenanceLogrepository struct {
	db *platform.Mysql
}

func NewMaintenanceLogRepository(db *platform.Mysql) domain.MaintenanceLogRepository {
	return &maintenanceLogrepository{db: db}
}

func (r *maintenanceLogrepository) GetAllmaintenanceLog() (*[]domain.MaintenanceLog, error) {
	response := new([]domain.MaintenanceLog)
	err := r.db.Select(response, "SELECT *"+
		" FROM `maintenanceLog`"+
		" INNER JOIN employee ON maintenanceLog.staffID = employee.employeeID")
	if err != nil {
		return nil, err
	}
	return response, nil
}
