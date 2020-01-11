package config

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// Config ally configuration
type Config struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Key   string         `gorm:"type:varchar(255);unique_index;not null" json:"key"`
	Value postgres.Jsonb `json:"value"`
}
