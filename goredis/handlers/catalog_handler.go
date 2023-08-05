package handlers

import (
	"goredis/services"

	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}

func NewcatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv: catalogSrv}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {

	products, err := h.catalogSrv.GetProduct()

	if err != nil {
		return err
	}

	response := fiber.Map{
		"status" : "ok",
		"products": products,
	}

	return c.JSON(response)

}
