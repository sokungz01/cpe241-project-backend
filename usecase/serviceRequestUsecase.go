package usecase

import (
	"fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type serviceRequestUsecase struct {
	serviceRepository domain.ServiceRequestRepository
	employee          domain.UserUseCase
	machine           domain.MachineUsecase
	errorType         domain.ErrorTypeUseCase
	errorlog          domain.ErrorlogRepository
}

func NewServiceRequestUsecase(serviceRepository domain.ServiceRequestRepository, employee domain.UserUseCase, machine domain.MachineUsecase, errorType domain.ErrorTypeUseCase, errorlog domain.ErrorlogRepository) domain.ServiceRequestUsecase {
	return &serviceRequestUsecase{
		serviceRepository: serviceRepository,
		employee:          employee,
		machine:           machine,
		errorType:         errorType,
		errorlog:          errorlog,
	}
}

func (u *serviceRequestUsecase) GetAllServiceRequest() (*[]domain.ServiceRequest, error) {
	response, err := u.serviceRepository.GetAllServiceRequest()
	if err != nil {
		return nil, err
	}
	fmt.Println(&response)
	return response, err
}

func (u *serviceRequestUsecase) CreateServiceRequest(newServiceRequest *domain.ServiceRequest) error {
	for _, v := range newServiceRequest.ErrorLog {
		u.errorlog.Create(v)
	}
	return nil
}
