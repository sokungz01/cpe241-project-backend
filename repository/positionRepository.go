package repository

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type positionRepository struct {
	db *platform.Mysql
}

func NewPositionRepository(db *platform.Mysql) domain.PositionRepository {
	return &positionRepository{db: db}
}

func (p *positionRepository) Create(position *domain.Position) error {
	_, err := p.db.NamedExec("INSERT INTO `position` (`positionName`,`positionSalary`)"+
		"VALUE (:positionname,:positionsalary)", position)
	if err != nil {
		return err
	}
	return nil
}

func (p *positionRepository) FindByPositionName(positionName string) (*domain.Position, error) {
	response := new(domain.Position)
	err := p.db.Get(response, "SELECT *"+
		"FROM `position`"+
		"WHERE `positionName` = ?", positionName)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *positionRepository) GetAll() (*[]domain.Position, error) {
	response := make([]domain.Position, 0)
	err := p.db.Select(&response, "SELECT *"+
		"FROM `position`")
	if err != nil {
		return nil, err
	}
	return &response, nil
}
