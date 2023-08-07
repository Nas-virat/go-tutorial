package main

import (
	_ "fmt"
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"
	_ "time"

	"github.com/go-redis/redis/v8"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogHandlerRedis(productService,redisClient)

	

	app := fiber.New()
	app.Get("/products",productHandler.GetProducts)
	app.Listen(":8000")

}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:root@tcp(127.0.0.1:3306)/fiber?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
