package repository

var roundConfig = resourceConfig{
	table:           "rounds",
	columns:         []string{"id", "event_id", "title", "round_number", "created_at"},
	writableColumns: []string{"id", "event_id", "title", "round_number", "created_at"},
	defaultOrderBy:  "round_number ASC",
}

func GetAllRounds() ([]map[string]any, error)                    { return getAll(roundConfig) }
func GetRoundByID(id string) (map[string]any, error)             { return getByID(roundConfig, id) }
func CreateRound(payload map[string]any) (map[string]any, error) { return create(roundConfig, payload) }
func UpdateRound(id string, payload map[string]any) (map[string]any, error) {
	return update(roundConfig, id, payload)
}
func DeleteRound(id string) error { return deleteByID(roundConfig, id) }
