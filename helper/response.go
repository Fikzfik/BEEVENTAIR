package helper

import "github.com/gofiber/fiber/v2"

func APIResponse(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func BadRequest(c *fiber.Ctx, message string) error {
	return errorResponse(c, fiber.StatusBadRequest, message)
}

func NotFound(c *fiber.Ctx, message string) error {
	return errorResponse(c, fiber.StatusNotFound, message)
}

func InternalError(c *fiber.Ctx, message string) error {
	return errorResponse(c, fiber.StatusInternalServerError, message)
}

func errorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}
