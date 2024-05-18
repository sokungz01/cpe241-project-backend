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

func (s *authenRepository) Me(userID int) (*domain.AuthenDetail, error) {
	var response domain.AuthenDetail
	err := s.db.Get(&response, "SELECT `employeeID`,`email`, `password`,`name`, `surname`, `positionID`, `imageURL` FROM `employee` WHERE `employeeID` = ? ", userID)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
