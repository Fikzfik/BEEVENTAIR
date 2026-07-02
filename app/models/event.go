package models

import "time"

type Event struct {
	ID                   string     `json:"id"`
	Slug                 string     `json:"slug"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Category             string     `json:"category"`
	Status               string     `json:"status"`
	StartDate            time.Time  `json:"start_date"`
	EndDate              time.Time  `json:"end_date"`
	Location             string     `json:"location"`
	ImageURL             *string    `json:"image_url"`
	PrizePool            string     `json:"prize_pool"`
	MaxTeams             int        `json:"max_teams"`
	CurrentTeams         int        `json:"current_teams"`
	RegistrationFee      float64    `json:"registration_fee"`
	EarlyBirdPrice       *float64   `json:"early_bird_price"`
	EarlyBirdQuota       *int       `json:"early_bird_quota"`
	RegistrationDeadline time.Time  `json:"registration_deadline"`
	TMDate               *time.Time `json:"tm_date"`
	OrganizerID          string     `json:"organizer_id"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}
