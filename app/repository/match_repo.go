package repository

var matchConfig = resourceConfig{
	table:           "matches",
	columns:         []string{"id", "event_id", "round_id", "team1_id", "team2_id", "team1_score", "team2_score", "winner_id", "status", "start_time", "created_at", "updated_at"},
	writableColumns: []string{"id", "event_id", "round_id", "team1_id", "team2_id", "team1_score", "team2_score", "winner_id", "status", "start_time", "created_at", "updated_at"},
	defaultOrderBy:  "created_at DESC",
	updatedAtColumn: "updated_at",
}

func GetAllMatches() ([]map[string]any, error)                   { return getAll(matchConfig) }
func GetMatchByID(id string) (map[string]any, error)             { return getByID(matchConfig, id) }
func CreateMatch(payload map[string]any) (map[string]any, error) { return create(matchConfig, payload) }
func UpdateMatch(id string, payload map[string]any) (map[string]any, error) {
	return update(matchConfig, id, payload)
}
func DeleteMatch(id string) error { return deleteByID(matchConfig, id) }
