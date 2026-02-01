package main

import (
	_ "fizzbuzz-v2/docs"
	"fizzbuzz-v2/internal/fizzbuzz"
	"fizzbuzz-v2/internal/health"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
)

// @title Fizzbuzz API
// @version 1.0
// @description Simple Fizzbuzz service
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Use(timeout.NewWithContext(func(context *fiber.Ctx) error {
		return context.Next()
	}, 5*time.Second))

	app.Get("/health", health.Handler)
	app.Get("/swagger/*", swagger.HandlerDefault)

	fizzBuzzRepository := fizzbuzz.NewInMemoryStorageRepository()
	fizzBuzzService := fizzbuzz.NewService(fizzBuzzRepository)
	fizzBuzzController := fizzbuzz.NewController(fizzBuzzService)
	fizzBuzzController.RegisterRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("listen error: %v", err)
	}
}
