package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetTeamMembers(c *fiber.Ctx) error { return list(c, repository.GetAllTeamMembers) }
func GetTeamMember(c *fiber.Ctx) error {
	return find(c, repository.GetTeamMemberByID, "Team member not found")
}
func CreateTeamMember(c *fiber.Ctx) error {
	return create(c, repository.CreateTeamMember, "Team member created successfully")
}
func UpdateTeamMember(c *fiber.Ctx) error {
	return update(c, repository.UpdateTeamMember, "Team member updated successfully", "Team member not found")
}
func DeleteTeamMember(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteTeamMember, "Team member deleted successfully", "Team member not found")
}
