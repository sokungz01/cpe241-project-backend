package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type authenRepository struct {
	db *platform.Mysql
}

func NewAuthenRepository(db *platform.Mysql) domain.AuthenRepository {
	return &authenRepository{db: db}
}

func (s *authenRepository) SignIn(user *domain.AuthenPayload) (*domain.AuthenPayload, error) {
	var response domain.AuthenPayload
	err := s.db.Get(&response, "SELECT `employeeID`,`email`, `password`, `positionID` FROM `employee` WHERE `email` = ? ", user.Email)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
