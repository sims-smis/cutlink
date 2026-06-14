package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sims-smis/cutlink.git/database"
	"github.com/sims-smis/cutlink.git/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/api/v1/links", routes.GetLinks)
}

func InitRedisClient() {
	database.RDB = database.ConnectClient(0)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env:", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New())

	setupRoutes(app)

	InitRedisClient()

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
