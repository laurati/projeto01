package domain

type Details struct {
	ID         string `gorm:"primary_key"`
	Version    int64
	BundleType string
}
