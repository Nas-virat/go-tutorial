package main

import (
	"fmt"
	_ "net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/gorilla/mux"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	//Middleware
	app.Use("/hello", func(c *fiber.Ctx) error {
		c.Locals("name", "bond") //local
		fmt.Println("before")
		err := c.Next()
		fmt.Println("after")
		return err
	})

	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	//GET
	//test by curl localhost:8000/hello -i
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Locals("names") //get local
		return c.SendString(fmt.Sprintf("GET: Hello World %v", name))
	})

	//POST
	//test by curl localhost:8000/hello -i -X POST
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("POST: Hello World")
	})

	//Parameter
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("name:" + name)
	})

	//Parameter Optional
	/*
		app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
			name := c.Params("name")
			surname := c.Params("surname")
			return c.SendString("Optional name:" + name + ",surname:" + surname)
		})
	*/
	//ParamsInt
	app.Get("/hello/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID = %v", id))
	})

	//Query
	//test by curl "localhost:8000/query?name=fang"
	//curl "localhost:8000/query?name=fang&surname=napas"
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name: " + name + "surname:" + surname)
	})

	//test by curl "localhost:8000/query2?id=1&name=fang"
	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	//Wildcars
	//test by curl localhost:8000/wildcards/hello/world
	//return hello/world
	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		wildcart := c.Params("*")
		return c.SendString(wildcart)
	})

	//Static file
	app.Static("/", "./wwwroot", fiber.Static{
		Index:         "index.html",
		CacheDuration: time.Second * 10,
	})

	//NewError
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	//Group
	// test by
	// curl localhost:8000/v1/hello
	v1 := app.Group("v1", func(c *fiber.Ctx) error{
		c.Set("Version", "v1")
		return c.Next()
	})


	v1.Get("/hello",func(c *fiber.Ctx) error {
		return c.SendString("Hello v1")
	})

	// test by
	// curl localhost:8000/v2/hello
	v2 := app.Group("v2", func(c *fiber.Ctx) error{
		c.Set("Version", "v1")
		return c.Next()
	})
	v2.Get("/hello",func(c *fiber.Ctx) error {
		return c.SendString("Hello v1")
	})

	//Mount
	

	app.Listen(":8000")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
