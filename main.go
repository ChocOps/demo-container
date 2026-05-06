package main

import (
	"math/rand/v2"
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

	/** cors */
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"service": os.Getenv("SERVICE"),
			"version": os.Getenv("VERSION"),
		})
	})

	app.Get("/motd", func(c *fiber.Ctx) error {
		var motds []string = []string{
			"Don't Panic.",
			"So long, and thanks for all the fish.",
			"Time is an illusion. Lunchtime doubly so.",
			"The Answer to Life, the Universe, and Everything is 42.",
			"We apologize for the inconvenience.",
			"Here I am, brain the size of a planet, and they ask me to serve HTTP requests.",
			"I'd far rather be happy than right any day.",
			"In the beginning, the Universe was created. This made a lot of people very angry.",
			"Life: quite pleasant if it weren't for all the people.",
			"The ships hung in the sky in much the same way that bricks don't.",
			"Nothing travels faster than light, except bad news.",
			"Vogons: not actually the worst poets in the Universe, just the third worst.",
			"A cup of tea would restore my normality.",
			"Space is big. Really big.",
			"It's at times like these I wish I had listened to what my mother told me. Why? What did she say? I don't know, I didn't listen.",
			"The major difference between a thing that might go wrong and a thing that cannot possibly go wrong is that when a thing that cannot possibly go wrong goes wrong it usually turns out to be impossible to get at or repair.",
			"He felt that his whole life was some kind of dream and he sometimes wondered whose it was.",
			"Anyone who is capable of getting themselves made president should on no account be allowed to do the job.",
			"Flying is learning how to throw yourself at the ground and miss.",
			"Resistance is useless!",
		}

		index := rand.IntN(len(motds))
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"motd": motds[index],
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
