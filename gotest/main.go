package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Fang", 21, nil)
	c.On("GetCustomer", 2).Return("Napas", 21, errors.New("Not Found"))
	c.On("GetCustomer", 3).Return("Nas-virat", 21, nil)

	//////////

	name, age, err := c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)
}

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
	Hello()
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (r *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := r.Called()
	return args.String(0), args.Int(1), args.Error(2)
}
