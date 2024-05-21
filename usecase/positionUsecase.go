package usecase

import (
	//"errors"
	//"fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type positionUsecase struct {
	positionRepo domain.PositionRepository
}

func NewPositionUsecase(positionRepo domain.PositionRepository) domain.PositionUsecase {
	return &positionUsecase{positionRepo: positionRepo}
}

func (pUse *positionUsecase) Create(position *domain.Position) error {
	if err := pUse.positionRepo.Create(position); err != nil {
		return err
	}
	return nil
}

func (pUse *positionUsecase) FindByPositionName(positionName string) (*domain.Position, error) {
	response, err := pUse.positionRepo.FindByPositionName(positionName)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (pUse *positionUsecase) GetPositionByID(id int) (*domain.Position, error) {
	response, err := pUse.positionRepo.GetPositionByID(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (pUse *positionUsecase) UpdatePosition(id int, newPositionData *domain.Position) error {
	updateErr := pUse.positionRepo.UpdatePosition(id, newPositionData)
	return updateErr
}

func (pUse *positionUsecase) DeletePosition(id int) error {
	return pUse.positionRepo.DeletePosition(id)
}

func (pUse *positionUsecase) GetAll() (*[]domain.Position, error) {
	response, err := pUse.positionRepo.GetAll()
	if err != nil {
		return nil, nil
	}
	return response, nil
}
