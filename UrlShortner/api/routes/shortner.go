package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/Kaushikmak/UrlShortner/db"
	"github.com/Kaushikmak/UrlShortner/helper"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// FIXED: Expiry is now a string to accept "24h", "10m", etc.
type request struct {
	URL            string `json:"url"`
	CustomShortner string `json:"customshortner"`
	Expiry         string `json:"expiry"`
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

	// --- Rate Limiting ---
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
	}

	// --- Validation ---
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	if helper.DomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Domain violation"})
	}

	body.URL = helper.EnforceHTTP(body.URL)

	// --- ID Generation ---
	var id string
	if body.CustomShortner == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShortner
	}

	redisDB_0 := db.CreateClient(0)
	defer redisDB_0.Close()

	//  Default to 24h if empty
	if body.Expiry == "" {
		body.Expiry = "24h"
	}

	// Parse the string (e.g., "24h") into a time.Duration object
	expiryDuration, err := time.ParseDuration(body.Expiry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expiry format. Use 24h, 10m, etc."})
	}

	// Use the parsed duration (expiryDuration) for Redis
	success, err := redisDB_0.SetNX(db.Ctx, id, body.URL, expiryDuration).Result()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to connect to server"})
	}

	if !success {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Custom shortner already in use"})
	}

	// --- Response ---
	resp := response{
		URL:             body.URL,
		CustomShortner:  "",
		Expiry:          expiryDuration,
		XRateRemaining:  10,
		XRateLimitReset: 30 * time.Minute,
	}

	redisDB_1.Decr(db.Ctx, c.IP())

	val, _ = redisDB_1.Get(db.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := redisDB_1.TTL(db.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShortner = os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
