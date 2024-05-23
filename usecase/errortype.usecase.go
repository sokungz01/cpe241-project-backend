package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type errorTypeUsecase struct {
	errorTypeRepository domain.ErrorTypeRepository
}

func NewErrorTypeUsecase(errorTypeRepository domain.ErrorTypeRepository) domain.ErrorTypeUseCase {
	return &errorTypeUsecase{errorTypeRepository: errorTypeRepository}
}

func (u *errorTypeUsecase) CreateErrorType(elem *domain.ErrorType) (*domain.ErrorType, error) {
	if elem.ErrorName == "" {
		return nil, errors.New("erorr! body empty")
	}
	response, err := u.errorTypeRepository.CreateErrorType(elem)
	if err != nil {
		return nil, errors.New("erorr! cannot create new error type")
	}
	return response, nil
}

func (u *errorTypeUsecase) GetAllErrorType() (*[]domain.ErrorType, error) {
	response, err := u.errorTypeRepository.GetAllErrorType()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *errorTypeUsecase) FindByID(id int) (*domain.ErrorType, error) {
	response, err := u.errorTypeRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *errorTypeUsecase) UpdateErrorType(id int, elem *domain.ErrorType) (*domain.ErrorType, error) {
	if elem.ErrorName == "" || id == 0 {
		return nil, errors.New("erorr! body empty")
	}
	response, err := u.errorTypeRepository.UpdateErrorType(id, elem)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *errorTypeUsecase) DeleteErrorType(id int) error {
	err := u.errorTypeRepository.DeleteErrorType(id)
	if err != nil {
		return err
	}
	return nil
}
