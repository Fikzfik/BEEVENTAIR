package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"eventbe/database"
)

var userConfig = resourceConfig{
	table:           "users",
	columns:         []string{"id", "username", "email", "password_hash", "role", "avatar_url", "created_at", "updated_at"},
	writableColumns: []string{"id", "username", "email", "password_hash", "role", "avatar_url", "created_at", "updated_at"},
	defaultOrderBy:  "created_at DESC",
	updatedAtColumn: "updated_at",
}

func GetAllUsers() ([]map[string]any, error)                    { return getAll(userConfig) }
func GetUserByID(id string) (map[string]any, error)             { return getByID(userConfig, id) }
func CreateUser(payload map[string]any) (map[string]any, error) { return create(userConfig, payload) }
func UpdateUser(id string, payload map[string]any) (map[string]any, error) {
	return update(userConfig, id, payload)
}
func DeleteUser(id string) error { return deleteByID(userConfig, id) }

func GetUserByEmail(email string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT %s FROM %s WHERE email = ?", strings.Join(userConfig.columns, ", "), userConfig.table)
	row := database.DB.QueryRowContext(ctx, query, email)

	return scanRow(row, userConfig.columns)
}

func GetUserByUsername(username string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ?", strings.Join(userConfig.columns, ", "), userConfig.table)
	row := database.DB.QueryRowContext(ctx, query, username)

	return scanRow(row, userConfig.columns)
}

