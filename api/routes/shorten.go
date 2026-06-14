package routes

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sims-smis/cutlink.git/database"
)

type Request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      int    `json:"expiry"`
}

type Response struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
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

	shortExists, _ := database.RDB.Get(
		database.Ctx,
		short,
	).Result()

	if shortExists != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "url custom short is already in use. Try with some other custom short.",
		})
	}

	erro := database.RDB.Set(
		database.Ctx,
		short,
		body.URL,
		time.Duration(body.Expiry)*time.Hour,
	).Err()

	fmt.Println("error is : ", erro)

	if erro != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not save URL",
		})
	}

	fmt.Println("Key added ", short)

	// return nil
	response := Response{
		URL:      body.URL,
		ShortURL: os.Getenv("DOMAIN") + "/" + short,
		Expiry:   body.Expiry,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
