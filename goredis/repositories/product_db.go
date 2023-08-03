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
	return productRepositoryDB{db: db}
}

func mockData(db *gorm.DB) error {

	//set once a time
	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	//random
	seed := rand.NewSource(time.Now().UnixNano())

	random := rand.New(seed)

	products := []product{}
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product%v", i+1),
			Quantity: random.Intn(100),
		})
	}
	return db.Create(&products).Error
}

func (r productRepositoryDB) GetProduct() (products []product, err error) {

	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products,err
}
