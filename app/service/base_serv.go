package service

import (
	"errors"
	"time"

	"eventbe/app/repository"
	"eventbe/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func list(c *fiber.Ctx, fetch func() ([]map[string]any, error)) error {
	data, err := fetch()
	if err != nil {
		return helper.InternalError(c, err.Error())
	}

	return helper.APIResponse(c, fiber.StatusOK, "Success", data)
}

func find(c *fiber.Ctx, fetch func(string) (map[string]any, error), notFoundMessage string) error {
	id := c.Params("id")
	if id == "" {
		return helper.BadRequest(c, "ID is required")
	}

	data, err := fetch(id)
	if err != nil {
		return respondRepositoryError(c, err, notFoundMessage)
	}

	return helper.APIResponse(c, fiber.StatusOK, "Success", data)
}

func create(c *fiber.Ctx, save func(map[string]any) (map[string]any, error), message string) error {
	input, err := parseInput(c)
	if err != nil {
		return helper.BadRequest(c, "Invalid input payload")
	}

	now := time.Now()
	input["id"] = uuid.New().String()
	setDefaultTimestamp(input, "created_at", now)
	setDefaultTimestamp(input, "updated_at", now)
	setDefaultTimestamp(input, "timestamp", now)
	setDefaultTimestamp(input, "joined_at", now)

	data, err := save(input)
	if err != nil {
		return helper.InternalError(c, err.Error())
	}

	return helper.APIResponse(c, fiber.StatusCreated, message, data)
}

func update(c *fiber.Ctx, save func(string, map[string]any) (map[string]any, error), message, notFoundMessage string) error {
	id := c.Params("id")
	if id == "" {
		return helper.BadRequest(c, "ID is required")
	}

	input, err := parseInput(c)
	if err != nil {
		return helper.BadRequest(c, "Invalid input payload")
	}
	delete(input, "id")

	data, err := save(id, input)
	if err != nil {
		return respondRepositoryError(c, err, notFoundMessage)
	}

	return helper.APIResponse(c, fiber.StatusOK, message, data)
}

func destroy(c *fiber.Ctx, remove func(string) error, message, notFoundMessage string) error {
	id := c.Params("id")
	if id == "" {
		return helper.BadRequest(c, "ID is required")
	}

	if err := remove(id); err != nil {
		return respondRepositoryError(c, err, notFoundMessage)
	}

	return helper.APIResponse(c, fiber.StatusOK, message, nil)
}

func parseInput(c *fiber.Ctx) (map[string]any, error) {
	var input map[string]any
	if err := c.BodyParser(&input); err != nil {
		return nil, err
	}
	if input == nil {
		input = map[string]any{}
	}

	return input, nil
}

func setDefaultTimestamp(input map[string]any, field string, value time.Time) {
	if _, ok := input[field]; !ok {
		input[field] = value
	}
}

func respondRepositoryError(c *fiber.Ctx, err error, notFoundMessage string) error {
	if errors.Is(err, repository.ErrNotFound) {
		return helper.NotFound(c, notFoundMessage)
	}
	return helper.InternalError(c, err.Error())
}
