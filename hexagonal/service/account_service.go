package service

import "bank/repository"

type accountService struct{
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService{
	return accountService{accRepo: accRepo}
}


func (s accountService) NewAccount(CustomerID int, NewAccountRequest) (*AccountResponse, error){

	// validate the request
	s.accRepo.Create(acc)
	return nil,nil
}

func (s accountService) GetAccounts(CustomerID int) ([]AccountResponse, error){
	accounts, err := s.accRepo.GetAll(CustomerID)
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