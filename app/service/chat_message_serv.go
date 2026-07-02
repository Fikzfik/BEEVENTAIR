package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetChatMessages(c *fiber.Ctx) error { return list(c, repository.GetAllChatMessages) }
func GetChatMessage(c *fiber.Ctx) error {
	return find(c, repository.GetChatMessageByID, "Chat message not found")
}
func CreateChatMessage(c *fiber.Ctx) error {
	return create(c, repository.CreateChatMessage, "Chat message created successfully")
}
func UpdateChatMessage(c *fiber.Ctx) error {
	return update(c, repository.UpdateChatMessage, "Chat message updated successfully", "Chat message not found")
}
func DeleteChatMessage(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteChatMessage, "Chat message deleted successfully", "Chat message not found")
}
