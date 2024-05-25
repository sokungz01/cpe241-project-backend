package usecase

import (
	"errors"

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

func (u *serviceResponseUsecase) CreateServiceResponse(newResponse *domain.ServiceResponse) error {
	_, userErr := u.user.GetById(newResponse.StaffID)
	if userErr != nil {
		return errors.New("serviceResponse: staffID error")
	}
	_, serviceErr := u.requestedService.GetServiceRequest(newResponse.RequestedServiceID)
	if serviceErr != nil {
		return errors.New("serviceResponnse: not a valid service request")
	}
	err := u.responseRepo.CreateServiceResponse(newResponse)
	return err
}

func (u *serviceResponseUsecase) DeleteResponse(id int) error {
	return u.responseRepo.DeleteResponse(id)
}

func (u *serviceResponseUsecase) GetResponse(id int) (*domain.ServiceResponse, error) {
	response, err := u.responseRepo.GetResponse(id)
	if err != nil {
		return nil, errors.New("serviceResponse: responseID error")
	}
	return response, err
}

func (u *serviceResponseUsecase) GetResponseByService(id int) (*[]domain.ServiceResponse, error) {
	response, err := u.responseRepo.GetResponseByService(id)
	if err != nil {
		return nil, errors.New("serviceResponse: serviceID error")
	}
	return response, err
}
