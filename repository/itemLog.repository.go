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

func (r *ItemLog) CreateItemLog(itemLog *domain.ItemLog) (*domain.ItemLog, error) {

	_, err := r.db.NamedExec("INSERT INTO `itemLog` (`itemID`, `qty`, `staffID`,`isAdd`)"+
		"VALUE (:itemID, :qty, :staffID, :isAdd)", itemLog)
	if err != nil {
		return nil, err
	}

	response := new(domain.ItemLog)
	_ = r.db.Get(response, "SELECT * FROM `itemLog` AS il INNER JOIN `inventory` AS item ON item.itemID = il.itemID INNER JOIN `employee` ON employee.employeeID = il.staffID WHERE `itemLogID` IN (SELECT LAST_INSERT_ID() as id)")
	return response, nil
}
