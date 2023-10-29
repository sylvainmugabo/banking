package services

import (
	"github.com/sylvainmugabo/microservices-lib/errs"
	"github.com/sylvainmugabo/microservices/banking/domain"
	"github.com/sylvainmugabo/microservices/banking/dto"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers := make([]dto.CustomerResponse, 0)
	customersDomain, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	for _, cust := range customersDomain {
		customers = append(customers, cust.ToDto())
	}

	return customers, nil

}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customerDomain, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}
	response := customerDomain.ToDto()
	return &response, nil

}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
