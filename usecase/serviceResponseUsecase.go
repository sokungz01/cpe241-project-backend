package usecase

import (
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type serviceResponseUsecase struct {
	responseRepo     domain.ServiceResponseRepository
	user             domain.UserUseCase
	requestedService domain.ServiceRequestUsecase
}

func NewServiceResponsUsecase(responseRepo domain.ServiceResponseRepository,
	user domain.UserUseCase, requestedService domain.ServiceRequestUsecase) domain.ServiceResponseUsecase {
	return &serviceResponseUsecase{
		responseRepo:     responseRepo,
		user:             user,
		requestedService: requestedService,
	}
}

func (u *serviceResponseUsecase) GetAllResponse() (*[]domain.ServiceResponse, error) {
	response, err := u.responseRepo.GetAllResponse()
	return response, err
}
