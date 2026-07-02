package repository

var prizeConfig = resourceConfig{
	table:           "prizes",
	columns:         []string{"id", "event_id", "rank", "reward"},
	writableColumns: []string{"id", "event_id", "rank", "reward"},
	defaultOrderBy:  "rank ASC",
}

func GetAllPrizes() ([]map[string]any, error)                    { return getAll(prizeConfig) }
func GetPrizeByID(id string) (map[string]any, error)             { return getByID(prizeConfig, id) }
func CreatePrize(payload map[string]any) (map[string]any, error) { return create(prizeConfig, payload) }
func UpdatePrize(id string, payload map[string]any) (map[string]any, error) {
	return update(prizeConfig, id, payload)
}
func DeletePrize(id string) error { return deleteByID(prizeConfig, id) }
