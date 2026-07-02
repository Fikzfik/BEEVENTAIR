package service

import (
	"errors"
	"time"

	"eventbe/app/repository"
	"eventbe/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest defines the input payload for user registration
type RegisterRequest struct {
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password_hash"` // maps to password_hash from frontend
	Role      string  `json:"role"`
	AvatarURL *string `json:"avatar_url"`
}

// LoginRequest defines the input payload for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register handles user registration
func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest(c, "Invalid input payload")
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return helper.BadRequest(c, "Username, email, and password are required")
	}

	// Validate role
	if req.Role == "" {
		req.Role = "participant"
	}
	if req.Role != "participant" && req.Role != "organizer" && req.Role != "admin" {
		return helper.BadRequest(c, "Invalid role. Role must be participant, organizer, or admin")
	}

	// Check if email already registered
	if _, err := repository.GetUserByEmail(req.Email); err == nil {
		return helper.BadRequest(c, "Email already registered")
	}

	// Check if username already taken
	if _, err := repository.GetUserByUsername(req.Username); err == nil {
		return helper.BadRequest(c, "Username already taken")
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return helper.InternalError(c, "Failed to hash password")
	}

	// Generate UUID and current timestamp
	id := uuid.New().String()
	now := time.Now()

	payload := map[string]any{
		"id":            id,
		"username":      req.Username,
		"email":         req.Email,
		"password_hash": string(hashedPassword),
		"role":          req.Role,
		"avatar_url":    req.AvatarURL,
		"created_at":    now,
		"updated_at":    now,
	}

	// Save to DB
	data, err := repository.CreateUser(payload)
	if err != nil {
		return helper.InternalError(c, err.Error())
	}

	// Remove password hash from response
	delete(data, "password_hash")

	// Map avatar_url to avatarUrl for frontend compatibility
	if avatarVal, ok := data["avatar_url"]; ok {
		data["avatarUrl"] = avatarVal
	}

	return helper.APIResponse(c, fiber.StatusCreated, "Registration successful", data)
}

// Login handles user authentication
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest(c, "Invalid input payload")
	}

	if req.Email == "" || req.Password == "" {
		return helper.BadRequest(c, "Email and password are required")
	}

	// Fetch user by email
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.NotFound(c, "Email not found")
		}
		return helper.InternalError(c, err.Error())
	}

	dbPasswordHash, ok := user["password_hash"].(string)
	if !ok {
		return helper.InternalError(c, "Invalid password stored in database")
	}

	// Verify password
	// 1. Try to compare using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(req.Password))
	if err != nil {
		// 2. Fallback: check plain text password (for legacy seed/mock data)
		if dbPasswordHash != req.Password {
			return helper.BadRequest(c, "Incorrect password")
		}
	}

	// Remove password hash from response
	delete(user, "password_hash")

	// Map avatar_url to avatarUrl for frontend compatibility
	if avatarVal, ok := user["avatar_url"]; ok {
		user["avatarUrl"] = avatarVal
	}

	// Generate a simple token response for the client (mock jwt token)
	// We return the user object alongside a mock token
	responsePayload := map[string]any{
		"id":        user["id"],
		"username":  user["username"],
		"email":     user["email"],
		"role":      user["role"],
		"avatarUrl": user["avatarUrl"],
		"token":     "mock-jwt-token-for-" + req.Email,
	}

	return helper.APIResponse(c, fiber.StatusOK, "Login successful", responsePayload)
}
