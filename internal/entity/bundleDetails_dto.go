package entity

import "time"

type DetailsDtoInput struct {
	Version    int64  `json:"version"`
	BundleType string `json:"bundle_type"`
	CreatedBy  string `json:"created_by"`
}

type DetailsDtoOutput struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
