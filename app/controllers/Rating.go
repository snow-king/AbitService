package Handlers

import (
	"AbitService/app/service"
	"AbitService/app/service/BrokerRepository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CompGroup(c *fiber.Ctx) error {
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
func RatingByGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Не верный id"})
	}
	rating, err := BrokerRepository.RatingList()
	if err != nil {
		return err
	}
	calc := service.NewRatingCalc(rating)
	list, err := calc.ByGroup(id)
	if err != nil {
		return err
	}
	return c.JSON(list)
}
func TopPriority(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Не верный id"})
	}
	rating, err := BrokerRepository.RatingList()
	if err != nil {
		return err
	}
	calc := service.NewRatingCalc(rating)
	group := calc.TopPriority(id)
	return c.JSON(fiber.Map{"group_id": group})
}
