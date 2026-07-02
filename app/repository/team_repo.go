package repository

var teamConfig = resourceConfig{
	table:           "teams",
	columns:         []string{"id", "event_id", "name", "logo_url", "captain_id", "created_at"},
	writableColumns: []string{"id", "event_id", "name", "logo_url", "captain_id", "created_at"},
	defaultOrderBy:  "created_at DESC",
}

func GetAllTeams() ([]map[string]any, error)                    { return getAll(teamConfig) }
func GetTeamByID(id string) (map[string]any, error)             { return getByID(teamConfig, id) }
func CreateTeam(payload map[string]any) (map[string]any, error) { return create(teamConfig, payload) }
func UpdateTeam(id string, payload map[string]any) (map[string]any, error) {
	return update(teamConfig, id, payload)
}
func DeleteTeam(id string) error { return deleteByID(teamConfig, id) }
