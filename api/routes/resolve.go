package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/sims-smis/cutlink.git/database"
)

func ResolveURL(c fiber.Ctx) error {
	fmt.Println("Inside Resolve URL Function")
	short := c.Params("url")

	url, err := database.RDB.Get(
		database.Ctx,
		short,
	).Result()

	if err != nil {
		fmt.Println("Error Occured")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short URL not found",
		})
	}
	fmt.Println("Got the result ", url)
	return c.Redirect().To(url)
}
