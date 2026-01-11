package routes

import (
	"github.com/Kaushikmak/UrlShortner/db"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c fiber.Ctx) error {
	url := c.Params("url")

	redisDBclient := db.CreateClient(0)
	defer redisDBclient.Close()

	value, err := redisDBclient.Get(db.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short url not found in database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to database"})
	}

	redisDBclient_Increment := db.CreateClient(1)
	defer redisDBclient_Increment.Close()

	_ = redisDBclient_Increment.Incr(db.Ctx, "counter")

	return c.Redirect().Status(fiber.StatusMovedPermanently).To(value)
}
