package repository

var chatMessageConfig = resourceConfig{
	table:           "chat_messages",
	columns:         []string{"id", "channel_name", "sender_id", "content", "timestamp"},
	writableColumns: []string{"id", "channel_name", "sender_id", "content", "timestamp"},
	defaultOrderBy:  "timestamp DESC",
}

func GetAllChatMessages() ([]map[string]any, error)        { return getAll(chatMessageConfig) }
func GetChatMessageByID(id string) (map[string]any, error) { return getByID(chatMessageConfig, id) }
func CreateChatMessage(payload map[string]any) (map[string]any, error) {
	return create(chatMessageConfig, payload)
}
func UpdateChatMessage(id string, payload map[string]any) (map[string]any, error) {
	return update(chatMessageConfig, id, payload)
}
func DeleteChatMessage(id string) error { return deleteByID(chatMessageConfig, id) }
