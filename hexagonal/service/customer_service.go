package service

import (
	"bank/repository"
	"database/sql"
	"errors"
	"log"
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
		log.Println(err)
		return nil, err
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

		if err == sql.ErrNoRows{
			return nil,errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}
	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerResponse, nil
}
