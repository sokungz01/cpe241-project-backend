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
	if err := r.db.Select(&response, "SELECT * FROM `itemLog`"); err != nil {
		return nil, err
	}
	return &response, nil
}
