package repository

import (
	"errors"
	"time"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "110011", DateOfBirth: date, Status: 1},
		{CustomerID: 1002, Name: "Rob", City: "New Delhi", ZipCode: "110011", DateOfBirth: date, Status: 0},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {

	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
