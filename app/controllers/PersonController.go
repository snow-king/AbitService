package PersonControllers

import (
	"AbitService/app/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Index(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Не верный id"})
	}
	person := new(service.PersonService)
	response := person.Show(id)
	return c.Status(fiber.StatusOK).JSON(response)
}
func ShowFamily(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Не верный id"})
	}
	person := new(service.PersonService)
	response := person.GetFamily(id)
	return c.Status(fiber.StatusOK).JSON(response)
}
