package domain

import "time"

type DetailsDtoInput struct {
	Version    int64
	BundleType string
	CreatedBy  string
}

type DetailsDtoOutput struct {
	ID        string
	Status    string
	CreatedBy string
	CreatedAt time.Time
}
