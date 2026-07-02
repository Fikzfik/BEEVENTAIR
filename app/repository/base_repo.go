package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"eventbe/database"
)

var ErrNotFound = errors.New("data not found")

type resourceConfig struct {
	table           string
	columns         []string
	writableColumns []string
	defaultOrderBy  string
	updatedAtColumn string
}

func getAll(config resourceConfig) ([]map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(config.columns, ", "), config.table)
	if config.defaultOrderBy != "" {
		query += " ORDER BY " + config.defaultOrderBy
	}

	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanRows(rows, config.columns)
}

func getByID(config resourceConfig, id string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", strings.Join(config.columns, ", "), config.table)
	row := database.DB.QueryRowContext(ctx, query, id)

	return scanRow(row, config.columns)
}

func create(config resourceConfig, payload map[string]any) (map[string]any, error) {
	columns, values := writablePayload(config, payload)
	if len(columns) == 0 {
		return nil, errors.New("request body does not contain writable fields")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	placeholders := make([]string, len(columns))
	for i := range columns {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		config.table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	if _, err := database.DB.ExecContext(ctx, query, values...); err != nil {
		return nil, err
	}

	id, ok := payload["id"].(string)
	if !ok || id == "" {
		return nil, errors.New("id is required")
	}

	return getByID(config, id)
}

func update(config resourceConfig, id string, payload map[string]any) (map[string]any, error) {
	columns, values := writablePayload(config, payload)
	if config.updatedAtColumn != "" {
		columns, values = removeColumn(columns, values, config.updatedAtColumn)
	}
	if len(columns) == 0 {
		return nil, errors.New("request body does not contain writable fields")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	setters := make([]string, len(columns))
	for i, column := range columns {
		setters[i] = column + " = ?"
	}
	if config.updatedAtColumn != "" {
		setters = append(setters, config.updatedAtColumn+" = CURRENT_TIMESTAMP")
	}

	values = append(values, id)
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = ?",
		config.table,
		strings.Join(setters, ", "),
	)

	result, err := database.DB.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, ErrNotFound
	}

	return getByID(config, id)
}

func deleteByID(config resourceConfig, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.DB.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = ?", config.table), id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}

	return nil
}

func writablePayload(config resourceConfig, payload map[string]any) ([]string, []any) {
	writable := make(map[string]bool, len(config.writableColumns))
	for _, column := range config.writableColumns {
		writable[column] = true
	}

	columns := make([]string, 0, len(payload))
	for column := range payload {
		if writable[column] {
			columns = append(columns, column)
		}
	}
	sort.Strings(columns)

	values := make([]any, len(columns))
	for i, column := range columns {
		values[i] = payload[column]
	}

	return columns, values
}

func removeColumn(columns []string, values []any, column string) ([]string, []any) {
	filteredColumns := make([]string, 0, len(columns))
	filteredValues := make([]any, 0, len(values))

	for i, current := range columns {
		if current == column {
			continue
		}
		filteredColumns = append(filteredColumns, current)
		filteredValues = append(filteredValues, values[i])
	}

	return filteredColumns, filteredValues
}

func scanRows(rows *sql.Rows, columns []string) ([]map[string]any, error) {
	items := []map[string]any{}
	for rows.Next() {
		item, err := scanScannable(rows, columns)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func scanRow(row *sql.Row, columns []string) (map[string]any, error) {
	item, err := scanScannable(row, columns)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return item, err
}

type scannable interface {
	Scan(dest ...any) error
}

func scanScannable(scanner scannable, columns []string) (map[string]any, error) {
	values := make([]any, len(columns))
	dest := make([]any, len(columns))
	for i := range values {
		dest[i] = &values[i]
	}

	if err := scanner.Scan(dest...); err != nil {
		return nil, err
	}

	item := make(map[string]any, len(columns))
	for i, column := range columns {
		if bytes, ok := values[i].([]byte); ok {
			item[column] = string(bytes)
			continue
		}
		item[column] = values[i]
	}

	return item, nil
}
