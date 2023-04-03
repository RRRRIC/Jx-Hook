package alertConfig

import (
	"time"

	"github.com/google/uuid"
)

type AlertConfig struct {
	ID           string            `json:"id"`
	SenderIds    map[string]string `json:"sender_ids,omitempty"`
	Enable       *bool             `json:"enable"`
	LastModified time.Time         `json:"last_modified,omitempty"`
}

type AlertSaveVO struct {
	ID        string            `path:"id" json:"id"`
	SenderIds map[string]string `json:"sender_ids,omitempty"`
	Enable    bool              `json:"enable"`
}

func (c *AlertSaveVO) ToConfig() AlertConfig {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return AlertConfig{
		ID:           c.ID,
		SenderIds:    c.SenderIds,
		Enable:       &c.Enable,
		LastModified: time.Now(),
	}
}
