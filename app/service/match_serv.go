package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetMatches(c *fiber.Ctx) error { return list(c, repository.GetAllMatches) }
func GetMatch(c *fiber.Ctx) error   { return find(c, repository.GetMatchByID, "Match not found") }
func CreateMatch(c *fiber.Ctx) error {
	return create(c, repository.CreateMatch, "Match created successfully")
}
func UpdateMatch(c *fiber.Ctx) error {
	return update(c, repository.UpdateMatch, "Match updated successfully", "Match not found")
}
func DeleteMatch(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteMatch, "Match deleted successfully", "Match not found")
}
