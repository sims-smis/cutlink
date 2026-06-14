package routes

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      int    `json:"expiry"`
}

type Response struct {
	URL      string `json:"url"`
	ShortURL string `json:"ShortURL"`
	Expiry   int    `json:"expiry"`
}

func ShortenURL(c fiber.Ctx) error {
	fmt.Println("Reached ShortenURL")

	body := new(Request)
	if err := c.Bind().Body(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate URL
	_, err := url.ParseRequestURI(body.URL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// Default expiry
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	// Generate short code if not provided
	short := body.CustomShort
	if short == "" {
		short = uuid.New().String()[:6]
	}
	// return nil
	response := Response{
		URL:      body.URL,
		ShortURL: os.Getenv("DOMAIN") + "/" + short,
		Expiry:   body.Expiry,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
