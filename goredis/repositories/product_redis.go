package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryRedis{db: db, redisClient: redisClient}
}

func (r productRepositoryRedis) GetProduct() (products []product, err error) {

	key := "repository::Getproducts"

	//Redis Get data from redis
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()

	//if data exist in redis
	if err == nil {
		//convert productsJson to object products
		err = json.Unmarshal([]byte(productsJson), &products)
		// if not err from convert to object
		if err == nil {
			fmt.Println("Get data from redis")
			return products, nil
		}
	}

	//database
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Redis ,Set data to redis

	data, err := json.Marshal(products) //convert object to json
	if err != nil {
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), 0).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("Get data from database")
	return products, nil
}
