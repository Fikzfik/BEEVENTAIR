package repository

var rundownItemConfig = resourceConfig{
	table:           "rundown_items",
	columns:         []string{"id", "event_id", "label", "start_time", "end_time"},
	writableColumns: []string{"id", "event_id", "label", "start_time", "end_time"},
	defaultOrderBy:  "start_time ASC",
}

func GetAllRundownItems() ([]map[string]any, error)        { return getAll(rundownItemConfig) }
func GetRundownItemByID(id string) (map[string]any, error) { return getByID(rundownItemConfig, id) }
func CreateRundownItem(payload map[string]any) (map[string]any, error) {
	return create(rundownItemConfig, payload)
}
func UpdateRundownItem(id string, payload map[string]any) (map[string]any, error) {
	return update(rundownItemConfig, id, payload)
}
func DeleteRundownItem(id string) error { return deleteByID(rundownItemConfig, id) }
