package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerTeamRoutes(api fiber.Router) {
	r := api.Group("/teams")

	r.Get("/", service.GetTeams)
	r.Get("/:id", service.GetTeam)
	r.Post("/", service.CreateTeam)
	r.Put("/:id", service.UpdateTeam)
	r.Patch("/:id", service.UpdateTeam)
	r.Delete("/:id", service.DeleteTeam)
}
