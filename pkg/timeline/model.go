package timeline

import "time"

// Timeline timeline item
type Timeline struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  time.Time    `gorm:"not null"`

	Provider string `gorm:"unique_index:uix_timelines_provider_type_custom_id;not null"`
	Type string `gorm:"unique_index:uix_timelines_provider_type_custom_id;not null"`
	CustomID string `gorm:"unique_index:uix_timelines_provider_type_custom_id;not null"`
}

// Timeline provider
const (
	ProviderWeiBo      = "weibo"
)

// Timeline type
const (
	TypeUser      = "user"
)
