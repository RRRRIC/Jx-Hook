package alertConfig

import (
	"time"

	"github.com/google/uuid"
)

type AlertConfig struct {
	ID           string            `json:"id"`
	SenderMap    map[string]string `json:"sender_map,omitempty"`
	Enable       *bool             `json:"enable"`
	LastModified time.Time         `json:"last_modified,omitempty"`
}

type AlertSaveVO struct {
	ID        string   `path:"id" json:"id"`
	SenderIds []string `json:"sender_ids,omitempty"`
	Enable    bool     `json:"enable"`
}

func (c *AlertSaveVO) ToConfig() AlertConfig {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return AlertConfig{
		ID:           c.ID,
		Enable:       &c.Enable,
		LastModified: time.Now(),
	}
}
