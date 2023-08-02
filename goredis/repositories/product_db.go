package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	return productRepositoryDB{db:db}
}

func mockData(db *gorm.DB) error{

	//random
	seed := rand.NewSource(time.Now().UnixNano())

	random := rand.New(seed)

	products := []product{}
	for i:=0; i< 5000;i++{
		products =  append(products,product{
			Name: fmt.Sprintf("Product%v", i+1),
			Quantity: ,
		})
	}
}

func (r productRepositoryDB) GetProduct() ([]product, error){
	return nil,nil
}