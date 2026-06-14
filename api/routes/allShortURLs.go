package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sims-smis/cutlink.git/database"
)

func GetLinks(c fiber.Ctx) error {
	links, err := database.RDB.SMembers(
		database.Ctx,
		"all_links",
	).Result()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "could not fetch links",
		})
	}

	return c.JSON(fiber.Map{
		"count": len(links),
		"links": links,
	})
}
