package Handlers

import (
	"AbitService/app/service"
	"AbitService/app/service/BrokerRepository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CompGroup(c *fiber.Ctx) error {
	//err := BrokerRepository.RequestRating()
	//if err != nil {
	//	return err
	//}
	rating, err := BrokerRepository.RatingList()
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"groups": rating})
}
func ListClaimed(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Не верный id"})
	}
	admission := service.NewApplicationAdmission(id)
	list := admission.List()
	return c.JSON(list)
}
