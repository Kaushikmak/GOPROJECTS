package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kaushikmak/UrlShortner/db"
	"github.com/Kaushikmak/UrlShortner/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables for local testing
	godotenv.Load()

	// Initialize the Singleton Redis Connection
	db.InitRedis()

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Route Registration
	api := app.Group("/api/v1")
	api.Post("/", routes.ShortnerURL)

	app.Get("/:url", routes.ResolveURL)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(app.Listen(port))
}
