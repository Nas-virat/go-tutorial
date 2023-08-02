package repositories

type product struct{
	ID 			int 
	Name 		string
	Quantity 	int
}

type ProductRepository interface{
	GetProduct() ([]product, error)
}


