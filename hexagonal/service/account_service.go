package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"time"
)

type accountService struct{
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService{
	return accountService{accRepo: accRepo}
}


func (s accountService) NewAccount(customerID int,request NewAccountRequest) (*AccountResponse, error){

	// validate the request
	account := repository.Account{
		CustomerID: customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount: request.Amount,
		Status: 1,
	}
	
	newAcc, err := s.accRepo.Create(account)
	if err != nil{
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// create response for send to client side
	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response,nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error){
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	accountsResponse := []AccountResponse{}

	for _, acc := range accounts {
		accountsResponse = append(accountsResponse, AccountResponse{
			AccountID:   acc.AccountID,
			OpeningDate: acc.OpeningDate,
			AccountType: acc.AccountType,
			Amount:      acc.Amount,
			Status:      acc.Status,
		})
	}

	return accountsResponse, nil
}