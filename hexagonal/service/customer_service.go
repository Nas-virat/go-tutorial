package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return customerService{customerRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponses := []CustomerResponse{}
	for _, v := range customers {
		custResponses = append(custResponses,
			CustomerResponse{
				CustomerID: v.CustomerID,
				Name:       v.Name,
				Status:     v.Status,
			})
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer,err := s.customerRepo.GetById(id)
	if err != nil{

		// if error is not critical don't have to log but send 
		// to other 
		if err == sql.ErrNoRows{
			return nil,errs.NewNotfoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerResponse, nil
}
