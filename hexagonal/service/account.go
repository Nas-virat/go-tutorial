package service


// this file create a port between client and business logic

// from repository\account.go:

// type Account struct {
// 	AccountID   int     `db:"account_id"`
// 	CustomerID  int     `db:"customer_id"`
// 	OpeningDate string  `db:"opening_date"`
// 	AccountType string  `db:"account_type"`
// 	Amount      float64 `db:"amount"`
// 	Status      int     `db:"status"`
// }

// object json from client
type NewAccountRequest struct{
	// don't need AccountID because it's auto increment
	// don't need CustomerID because it from params /customer/{customer_id}/account
	// don't need OpeningDate because server should handle it because 
	// it's not good idea to let client handle it such as timezone problem
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	// don't need Status because it's default 1
}

// object json to client
type AccountResponse struct {
	AccountID   int     `json:"account_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      int     `json:"status"`
}

type AccountService interface{
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts(int) ([]AccountResponse, error)
}