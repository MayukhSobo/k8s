package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"msg": "pong",
	})
}


func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(
		logger.Config{
			Format:     "[LOG] ${time} ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "2006/01/2 15:04:05",
		},
	))
	app.Get("/ping", ping)
	
	PORT := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}