package Handlers

import (
	"AbitService/app/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type FamilyAdd struct {
	Token    string `validate:"required"`
	ParentId int
}

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
func AppendFamily(c *fiber.Ctx) error {
	request := new(FamilyAdd)
	err := c.BodyParser(request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "не верные данные "})
	}
	family := service.NewFamily(request.ParentId)
	err = family.GetAccess(request.Token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "не верные данные "})
	}
	return c.JSON(fiber.Map{"message": "Запрос на создание семьи отправлен"})
}
