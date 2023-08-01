package handlers_test

import (
	"errors"
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		//Arrange
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})
	//Error on ConvertString
	t.Run("strconv.Atoi", func(t *testing.T) {

		//Arrange
		amount := "{a:10}"

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0, errors.New(""))

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)

	})

	//Error on CaluculateDiscount
	t.Run("Error CaluculateDiscount", func(t *testing.T) {

		//Arrange
		amount := 100

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0, errors.New(""))

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)

	})
}
