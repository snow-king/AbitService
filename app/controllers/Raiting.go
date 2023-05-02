package Handlers

import (
	"AbitService/app/service"
	"github.com/gofiber/fiber/v2"
)

func CompGroup(c *fiber.Ctx) error {
	groups := service.GetGroups(1)
	return c.JSON(fiber.Map{"groups": groups})
}
