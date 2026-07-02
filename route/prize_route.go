package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerPrizeRoutes(api fiber.Router) {
	r := api.Group("/prizes")

	r.Get("/", service.GetPrizes)
	r.Get("/:id", service.GetPrize)
	r.Post("/", service.CreatePrize)
	r.Put("/:id", service.UpdatePrize)
	r.Patch("/:id", service.UpdatePrize)
	r.Delete("/:id", service.DeletePrize)
}
