package config

import (
	f "github.com/gofiber/fiber/v2"
)

func NewApp() *f.App {
	return f.New(f.Config{
		AppName: "EventNear API",
		BodyLimit: 50 * 1024 * 1024, // 50MB limit for base64 image payloads
	})
}
