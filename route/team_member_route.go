package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerTeamMemberRoutes(api fiber.Router) {
	r := api.Group("/team-members")

	r.Get("/", service.GetTeamMembers)
	r.Get("/:id", service.GetTeamMember)
	r.Post("/", service.CreateTeamMember)
	r.Put("/:id", service.UpdateTeamMember)
	r.Patch("/:id", service.UpdateTeamMember)
	r.Delete("/:id", service.DeleteTeamMember)
}
