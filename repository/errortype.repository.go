package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type errorTypeRepository struct {
	db *platform.Mysql
}

func NewErrorTypeRepository(db *platform.Mysql) domain.ErrorTypeRepository {
	return &errorTypeRepository{db: db}
}

func (r *errorTypeRepository) CreateErrorType(elem *domain.ErrorType) (*domain.ErrorType, error) {
	_, err := r.db.NamedExec("INSERT INTO `errorType` (`errorName`)"+
		"VALUE (:errorName)", elem)

	if err != nil {
		return nil, err
	}

	response := new(domain.ErrorType)
	_ = r.db.Get(response, "SELECT * FROM `errorType` WHERE errorTypeID IN (SELECT LAST_INSERT_ID() as id)")

	return response, nil
}

func (r *errorTypeRepository) GetAllErrorType() (*[]domain.ErrorType, error) {
	response := make([]domain.ErrorType, 0)
	if err := r.db.Select(&response, "SELECT * FROM `errorType`"); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *errorTypeRepository) FindByID(id int) (*domain.ErrorType, error) {
	response := new(domain.ErrorType)
	err := r.db.Get(response, "SELECT *"+
		"FROM `errorType`"+
		"WHERE `errorTypeID` = ?", id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *errorTypeRepository) UpdateErrorType(id int, elem *domain.ErrorType) (*domain.ErrorType, error) {
	_, err := r.db.Exec("UPDATE `errorType`"+
		"SET `errorName`= ? WHERE `errorTypeID`= ?", elem.ErrorName, id)
	if err != nil {
		return nil, err
	}
	response, _ := r.FindByID(id)
	return response, nil
}

func (r *errorTypeRepository) DeleteErrorType(id int) error {
	_, err := r.db.Exec("DELETE FROM `errorType` WHERE `errorTypeID` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
