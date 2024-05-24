package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type errorLogRepository struct {
	db *platform.Mysql
}

func NewErrorLogRepository(db *platform.Mysql) domain.ErrorlogRepository {
	return &errorLogRepository{
		db: db,
	}
}

func (elog *errorLogRepository) Create(newError domain.ErrorLog) error {
	_, insertErr := elog.db.Exec("INSERT INTO `errorLog` (`errorTypeID`,`serviceID`,`errorDescription`,`createdDate`,`updateDate`)"+
		"VALUE (:errorTypeID,:serviceID,:errorDescription,:createdDate,:updateDate)", newError)
	return insertErr
}
