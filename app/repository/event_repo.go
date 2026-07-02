package repository

var eventConfig = resourceConfig{
	table:           "events",
	columns:         []string{"id", "slug", "title", "description", "category", "status", "start_date", "end_date", "location", "image_url", "prize_pool", "max_teams", "current_teams", "registration_fee", "early_bird_price", "early_bird_quota", "registration_deadline", "tm_date", "organizer_id", "created_at", "updated_at"},
	writableColumns: []string{"id", "slug", "title", "description", "category", "status", "start_date", "end_date", "location", "image_url", "prize_pool", "max_teams", "current_teams", "registration_fee", "early_bird_price", "early_bird_quota", "registration_deadline", "tm_date", "organizer_id", "created_at", "updated_at"},
	defaultOrderBy:  "created_at DESC",
	updatedAtColumn: "updated_at",
}

func GetAllEvents() ([]map[string]any, error)                    { return getAll(eventConfig) }
func GetEventByID(id string) (map[string]any, error)             { return getByID(eventConfig, id) }
func CreateEvent(payload map[string]any) (map[string]any, error) { return create(eventConfig, payload) }
func UpdateEvent(id string, payload map[string]any) (map[string]any, error) {
	return update(eventConfig, id, payload)
}
func DeleteEvent(id string) error { return deleteByID(eventConfig, id) }
