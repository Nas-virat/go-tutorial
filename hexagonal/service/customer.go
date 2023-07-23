package service

// adapter
type CustomerResponse struct{
	CustomerID 	int		`json:"customer_id"`
	Name		string	`json:"name"`
	Status		int		`json:"status"`
}

// port 
type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}



