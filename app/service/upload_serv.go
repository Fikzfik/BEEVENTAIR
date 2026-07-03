package service

import (
	"eventbe/helper"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return helper.BadRequest(c, "No file uploaded or file field must be named 'file'")
	}

	// Create upload folder if not exists
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return helper.InternalError(c, "Failed to create upload directory: "+err.Error())
	}

	// Generate unique file name
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String(), ext)
	filePath := filepath.Join(uploadDir, newFileName)

	// Save file
	if err := c.SaveFile(file, filePath); err != nil {
		return helper.InternalError(c, "Failed to save file: "+err.Error())
	}

	// Return file path URL
	// e.g. http://localhost:5000/uploads/file.png or http://127.0.0.1:5000/uploads/file.png
	protocol := "http"
	if c.Secure() {
		protocol = "https"
	}
	
	// Construct hosted URL path
	fileURL := fmt.Sprintf("%s://%s/uploads/%s", protocol, c.Hostname(), newFileName)

	return helper.APIResponse(c, fiber.StatusOK, "File uploaded successfully", fileURL)
}
