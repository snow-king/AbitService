package router

import (
	PersonControllers "AbitService/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(router fiber.Router) {
	router.Get("/person/:id", PersonControllers.Index)
	router.Get("/person/:id/family", PersonControllers.ShowFamily)
}
