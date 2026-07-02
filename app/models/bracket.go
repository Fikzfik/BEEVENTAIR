package models

import "time"

type Round struct {
	ID          string    `json:"id"`
	EventID     string    `json:"event_id"`
	Title       string    `json:"title"`
	RoundNumber int       `json:"round_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type Match struct {
	ID         string     `json:"id"`
	EventID    string     `json:"event_id"`
	RoundID    string     `json:"round_id"`
	Team1ID    *string    `json:"team1_id"`
	Team2ID    *string    `json:"team2_id"`
	Team1Score int        `json:"team1_score"`
	Team2Score int        `json:"team2_score"`
	WinnerID   *string    `json:"winner_id"`
	Status     string     `json:"status"`
	StartTime  *time.Time `json:"start_time"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
