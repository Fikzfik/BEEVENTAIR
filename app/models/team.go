package models

import "time"

type Team struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	Name      string    `json:"name"`
	LogoURL   *string   `json:"logo_url"`
	CaptainID string    `json:"captain_id"`
	CreatedAt time.Time `json:"created_at"`
}

type TeamMember struct {
	ID       string    `json:"id"`
	TeamID   string    `json:"team_id"`
	UserID   string    `json:"user_id"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}
