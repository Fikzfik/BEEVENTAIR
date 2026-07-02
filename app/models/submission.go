package models

import "time"

type Submission struct {
	ID          string    `json:"id"`
	EventID     string    `json:"event_id"`
	TeamID      string    `json:"team_id"`
	SubmittedBy string    `json:"submitted_by"`
	FileURL     string    `json:"file_url"`
	FileName    string    `json:"file_name"`
	FileSize    int       `json:"file_size"`
	Status      string    `json:"status"`
	Feedback    *string   `json:"feedback"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
