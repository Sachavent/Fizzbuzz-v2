package health

import "github.com/gofiber/fiber/v2"

type HealthResponse struct {
	Status string `json:"status"`
}

// Handler godoc
// @Summary Health check
// @Description Returns service status
// @Tags health
// @Success 200 {object} HealthResponse
// @Router /health [get]
func Handler(context *fiber.Ctx) error {
	return context.JSON(HealthResponse{Status: "ok"})
}
