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
	router.Get("/person/:id/list", handlers.ListClaimed)
	router.Get("/group/:id/rating", handlers.RatingByGroup)
	router.Get("/health", func(context *fiber.Ctx) error {
		return context.JSON(fiber.Map{"response": "It's Alive! Alive!!!!"})
	})
	router.Get("/person/:id/TopPriority", handlers.TopPriority)
}
