package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetTeams(c *fiber.Ctx) error { return list(c, repository.GetAllTeams) }
func GetTeam(c *fiber.Ctx) error  { return find(c, repository.GetTeamByID, "Team not found") }
func CreateTeam(c *fiber.Ctx) error {
	return create(c, repository.CreateTeam, "Team created successfully")
}
func UpdateTeam(c *fiber.Ctx) error {
	return update(c, repository.UpdateTeam, "Team updated successfully", "Team not found")
}
func DeleteTeam(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteTeam, "Team deleted successfully", "Team not found")
}
