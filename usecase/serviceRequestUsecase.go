package usecase

import (
	"errors"

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
	for index, item := range *response {
		errorLog, err := u.errorlog.FindByServiceID(item.ServiceID)
		if err != nil {
			return nil, err
		}
		(*response)[index].ErrorLog = make([]domain.ErrorLog, 0)
		(*response)[index].ErrorLog = *errorLog
	}
	return response, err
}

func (u *serviceRequestUsecase) GetServiceRequest(id int) (*domain.ServiceRequest, error) {
	if id == 0 {
		return nil, errors.New("error! serviceID cannot be not a number ")
	}
	response, err := u.serviceRepository.GetServiceRequest(id)
	if err != nil {
		return nil, err
	}
	errorLog, err := u.errorlog.FindByServiceID(response.ServiceID)
	if err != nil {
		return nil, err
	}
	response.ErrorLog = make([]domain.ErrorLog, 0)
	response.ErrorLog = *errorLog
	return response, err
}

func (u *serviceRequestUsecase) CreateServiceRequest(newServiceRequest *domain.ServiceRequest) (*domain.ServiceRequest, error) {
	if newServiceRequest == nil || newServiceRequest.EmployeeID == 0 || newServiceRequest.MachineID == 0 || newServiceRequest.Description == "" {
		return nil, errors.New("error! service request data not provide")
	}

	response, err := u.serviceRepository.CreateServiceRequest(newServiceRequest)

	if err != nil {
		return nil, errors.New("error! service request cannot create")
	}

	serviceID := response.ServiceID
	errorLogArr := make([]domain.ErrorLog, 0)
	for _, v := range newServiceRequest.ErrorLog {
		v.ServiceID = serviceID
		response, err := u.errorlog.Create(&v)
		if err != nil {
			return nil, errors.New("error! errorlog cannot create")
		}
		errorLogArr = append(errorLogArr, *response)
	}

	response.ErrorLog = make([]domain.ErrorLog, 0)
	response.ErrorLog = errorLogArr
	return nil, nil
}

func (u *serviceRequestUsecase) UpdateServiceRequestStatus(id int, statusID int) (*domain.ServiceRequest, error) {

	statusIDList := []int{1, 2, 3, 4, 5, 6}
	isContains := false
	for _, elem := range statusIDList {
		if elem == statusID {
			isContains = true
			break
		}
	}
	if !isContains {
		return nil, errors.New("error! statusid is invalid")
	}

	if id == 0 {
		return nil, errors.New("error! id is not invalid")
	}

	_, err := u.serviceRepository.GetServiceRequest(id)

	if err != nil {
		return nil, errors.New("error! id is not contain in db")
	}
	response, err := u.serviceRepository.UpdateServiceRequestStatus(id, statusID)

	if err != nil {
		return nil, errors.New("error! service request cannot update")
	}

	return response, nil
}
