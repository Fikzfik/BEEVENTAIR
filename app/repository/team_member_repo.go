package repository

var teamMemberConfig = resourceConfig{
	table:           "team_members",
	columns:         []string{"id", "team_id", "user_id", "role", "joined_at"},
	writableColumns: []string{"id", "team_id", "user_id", "role", "joined_at"},
	defaultOrderBy:  "joined_at DESC",
}

func GetAllTeamMembers() ([]map[string]any, error)        { return getAll(teamMemberConfig) }
func GetTeamMemberByID(id string) (map[string]any, error) { return getByID(teamMemberConfig, id) }
func CreateTeamMember(payload map[string]any) (map[string]any, error) {
	return create(teamMemberConfig, payload)
}
func UpdateTeamMember(id string, payload map[string]any) (map[string]any, error) {
	return update(teamMemberConfig, id, payload)
}
func DeleteTeamMember(id string) error { return deleteByID(teamMemberConfig, id) }
