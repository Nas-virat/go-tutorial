package main

import (
	"fmt"
	"goredis/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()

	productRepo := repositories.NewProductRepository(db)
	products,err := productRepo.GetProduct()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(products)

	/*
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Millisecond * 10)
		return c.SendString("Hello World")
	})

	app.Listen(":8000")
	*/
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:root@tcp(127.0.0.1:3306)/fiber?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
