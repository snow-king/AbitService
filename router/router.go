package router

import (
	handlers "AbitService/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(router fiber.Router) {
	router.Get("/person/:id", handlers.Index)
	router.Get("/person/:id/family", handlers.ShowFamily)
	router.Get("/groups", handlers.CompGroup)
	router.Post("/family/create", handlers.AppendFamily)

}
