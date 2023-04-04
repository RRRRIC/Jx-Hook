package senderConfig

import (
	"time"

	"github.com/google/uuid"
)

type SenderConfig struct {
	ID             string    `json:"id"`
	Name           string    `json:"name,omitempty"`
	TemplateMsg    string    `json:"template_msg,omitempty"`
	WechatRobotKey string    `json:"wechat_robot_key,omitempty"`
	CustomUrl      string    `json:"custom_url,omitempty"`
	AlertType      string    `json:"alert_type,omitempty"`
	Enable         *bool     `json:"enable"`
	LastModified   time.Time `json:"last_modified,omitempty"`
}

type SenderSaveVO struct {
	ID             string `path:"id" json:"id"`
	Name           string `json:"name,omitempty"`
	TemplateMsg    string `json:"template_msg,omitempty"`
	WechatRobotKey string `json:"wechat_robot_key,omitempty" vd:"len($) > 10"`
	AlertType      string `json:"alert_type,omitempty"`
	CustomUrl      string `json:"custom_url,omitempty"`
	Enable         bool   `json:"enable"`
}

func (c *SenderSaveVO) ToConfig() SenderConfig {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return SenderConfig{
		ID:             c.ID,
		Name:           c.Name,
		TemplateMsg:    c.TemplateMsg,
		WechatRobotKey: c.WechatRobotKey,
		AlertType:      c.AlertType,
		CustomUrl:      c.CustomUrl,
		Enable:         &c.Enable,
		LastModified:   time.Now(),
	}
}
