package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func ResolveURL(c fiber.Ctx) error {
	short := c.Params("url")
	fmt.Println(short, " hehe ")

	return nil
}
