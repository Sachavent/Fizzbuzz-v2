package fizzbuzz

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) RegisterRoutes(app *fiber.App) {
	app.Get("/fizzbuzz/result", controller.GetResultHandler)
	app.Get("/fizzbuzz/stats", controller.MostRequestedHandler)
}

// GetResultHandler godoc
// @Summary Compute fizzbuzz
// @Description Returns fizzbuzz sequence based on query parameters
// @Tags fizzbuzz
// @Produce plain
// @Param int1 query int true "First divisor" minimum(1)
// @Param int2 query int true "Second divisor" minimum(1)
// @Param limit query int true "Upper limit" minimum(1)
// @Param str1 query string true "Replacement for multiples of int1"
// @Param str2 query string true "Replacement for multiples of int2"
// @Success 200 {string} string
// @Failure 400 {object} ErrorResponse
// @Router /fizzbuzz/result [get]
func (controller *Controller) GetResultHandler(context *fiber.Ctx) error {
	var query GetResultQuery
	if err := context.QueryParser(&query); err != nil {
		return context.Status(400).JSON(fiber.Map{
			"error": "invalid query parameters",
		})
	}

	validate := validator.New()
	if err := validate.Struct(query); err != nil {
		return context.Status(400).JSON(fiber.Map{
			"errors": strings.Split(err.Error(), "\n"),
		})
	}

	return context.SendString(controller.service.GetResult(query))
}

// MostRequestedHandler godoc
// @Summary Most requested parameters
// @Description Returns the most frequent fizzbuzz parameters and count
// @Tags fizzbuzz
// @Success 200 {object} GetMostFrequentRequestOutput
// @Router /fizzbuzz/stats [get]
func (controller *Controller) MostRequestedHandler(context *fiber.Ctx) error {
	params, count := controller.service.GetMostFrequentRequest()
	if count == 0 {
		return context.Status(400).JSON(fiber.Map{
			"error": "fizzbuzz endpoint has not been called yet",
		})
	}

	output := GetMostFrequentRequestOutput{
		Parameters: params,
		Count:      count,
	}

	return context.JSON(output)
}
