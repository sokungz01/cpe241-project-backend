package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type ItemLog struct {
	db *platform.Mysql
}

func NewItemLogRepository(db *platform.Mysql) domain.ItemLogRepository {
	return &ItemLog{db: db}
}

func (r *ItemLog) GetAll() (*[]domain.ItemLog, error) {
	response := make([]domain.ItemLog, 0)
	if err := r.db.Select(&response, "SELECT * FROM `itemLog` AS il INNER JOIN `inventory` AS item ON item.itemID = il.itemID INNER JOIN `employee` ON employee.employeeID = il.staffID"); err != nil {
		return nil, err
	}
	for i := range response {
		response[i].Staff.Password = ""
	}
	return &response, nil
}
