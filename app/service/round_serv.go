package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetRounds(c *fiber.Ctx) error { return list(c, repository.GetAllRounds) }
func GetRound(c *fiber.Ctx) error  { return find(c, repository.GetRoundByID, "Round not found") }
func CreateRound(c *fiber.Ctx) error {
	return create(c, repository.CreateRound, "Round created successfully")
}
func UpdateRound(c *fiber.Ctx) error {
	return update(c, repository.UpdateRound, "Round updated successfully", "Round not found")
}
func DeleteRound(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteRound, "Round deleted successfully", "Round not found")
}
