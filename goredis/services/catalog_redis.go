package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{productRepo: productRepo, redisClient: redisClient}
}

func (s catalogServiceRedis) GetProduct() (products []Product, err error) {

	key := "service::Getproducts"

	//Redis Read Data
	productsJson, err := s.redisClient.Get(context.Background(), key).Result()

	//can read data from redis
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		// if can convert Json -> Go Object
		if err == nil {
			fmt.Println("The data is from Redis")
			return products, nil
		}
	}

	//Repository
	productsDB, err := s.productRepo.GetProduct()

	if err != nil {
		return nil, err
	}

	for _, p := range productsDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// Redis SET

	// Go Object -> Json
	data, err := json.Marshal(products)

	if err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}


fmt.Println("database")
	return products, nil
}
