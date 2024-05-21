package usecase

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type itemLogUsecase struct {
	repository domain.ItemLogRepository
}

func NewItemLogUsecase(repository domain.ItemLogRepository) domain.ItemLogUsecase {
	return &itemLogUsecase{repository: repository}
}

func (u *itemLogUsecase) GetAll() (*[]domain.ItemLog, error) {
	response, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}
