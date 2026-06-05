package domain

import "time"

type Snapshot struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	Data      map[string]any `json:"data"`
}
