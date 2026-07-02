package repository

var submissionConfig = resourceConfig{
	table:           "submissions",
	columns:         []string{"id", "event_id", "team_id", "submitted_by", "file_url", "file_name", "file_size", "status", "feedback", "created_at", "updated_at"},
	writableColumns: []string{"id", "event_id", "team_id", "submitted_by", "file_url", "file_name", "file_size", "status", "feedback", "created_at", "updated_at"},
	defaultOrderBy:  "created_at DESC",
	updatedAtColumn: "updated_at",
}

func GetAllSubmissions() ([]map[string]any, error)        { return getAll(submissionConfig) }
func GetSubmissionByID(id string) (map[string]any, error) { return getByID(submissionConfig, id) }
func CreateSubmission(payload map[string]any) (map[string]any, error) {
	return create(submissionConfig, payload)
}
func UpdateSubmission(id string, payload map[string]any) (map[string]any, error) {
	return update(submissionConfig, id, payload)
}
func DeleteSubmission(id string) error { return deleteByID(submissionConfig, id) }
