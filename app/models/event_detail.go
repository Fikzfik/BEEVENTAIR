package models

import "time"

type Prize struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	Rank    string `json:"rank"`
	Reward  string `json:"reward"`
}

type RundownItem struct {
	ID        string     `json:"id"`
	EventID   string     `json:"event_id"`
	Label     string     `json:"label"`
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}
