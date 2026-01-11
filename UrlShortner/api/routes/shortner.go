package routes

import (
	"time"

	"github.com/Kaushikmak/UrlShortner/helper"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v3"
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
