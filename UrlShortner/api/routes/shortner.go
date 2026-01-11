package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/Kaushikmak/UrlShortner/db"
	"github.com/Kaushikmak/UrlShortner/helper"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type request struct {
	URL            string        `json:"url"`
	CustomShortner string        `json:"customshortner"`
	Expiry         time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShortner  string        `json:"customshort"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortnerURL(c fiber.Ctx) error {
	body := new(request)

	if err := c.Bind().JSON(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// rate limiting
	redisDB_1 := db.CreateClient(1)
	defer redisDB_1.Close()
	val, err := redisDB_1.Get(db.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = redisDB_1.Set(db.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := redisDB_1.TTL(db.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Second,
			})
		}
		redisDB_1.Decr(db.Ctx, c.IP())

	}

	// validate url
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for self looping domain
	if helper.DomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Domain violation"})
	}

	// enforce to use HTTP
	body.URL = helper.EnforceHTTP(body.URL)

	return nil
}
