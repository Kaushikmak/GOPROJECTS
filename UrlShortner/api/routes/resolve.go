package routes

import (
	"github.com/Kaushikmak/UrlShortner/db"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c fiber.Ctx) error {
	url := c.Params("url")

	// Retrieve the original URL using the global database client
	value, err := db.Client.Get(db.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short url not found in database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to database"})
	}

	// Increment the global click counter using the same connection pool
	_ = db.Client.Incr(db.Ctx, "counter").Err()

	// Execute HTTP 301 Redirect to the target destination
	return c.Redirect().Status(fiber.StatusMovedPermanently).To(value)
}
