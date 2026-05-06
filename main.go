package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
)

func main() {
	/** local development */
	if "" == os.Getenv("ENV") {
		godotenv.Load()
	}

	/** logger */
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
		os.Exit(1)
	}

	if os.Getenv("ENV") != "production" {
		logger = prettyconsole.NewLogger(zap.DebugLevel)
	}

	logger = logger.With(
		zap.String("service", os.Getenv("SERVICE")),
		zap.String("version", os.Getenv("VERSION")),
	)

	/** fiber */
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"service": os.Getenv("SERVICE"),
			"version": os.Getenv("VERSION"),
		})
	})

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	logger.Info("Listening", zap.String("port", port))
	if err := app.Listen(":" + port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
