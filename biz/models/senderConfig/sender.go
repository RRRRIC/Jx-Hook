package senderConfig

import (
	"time"

	"github.com/google/uuid"
)

type SenderConfig struct {
	ID              string    `json:"id"`
	Name            string    `json:"name,omitempty"`
	TemplateMsg     string    `json:"template_msg,omitempty"`
	WechatRobotKey  string    `json:"wechat_robot_key,omitempty"`
	WechatAlertType string    `json:"wechat_alert_type,omitempty"`
	Enable          *bool     `json:"enable"`
	LastModified    time.Time `json:"last_modified,omitempty"`
}

type SenderSaveVO struct {
	ID              string `path:"id" json:"id"`
	Name            string `json:"name,omitempty"`
	TemplateMsg     string `json:"template_msg,omitempty"`
	WechatRobotKey  string `json:"wechat_robot_key,omitempty" vd:"len($) > 10"`
	WechatAlertType string `default:"markdown,omitempty" json:"wechat_alert_type"`
	Enable          bool   `json:"enable"`
}

func (c *SenderSaveVO) ToConfig() SenderConfig {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return SenderConfig{
		ID:              c.ID,
		Name:            c.Name,
		TemplateMsg:     c.TemplateMsg,
		WechatRobotKey:  c.WechatRobotKey,
		WechatAlertType: c.WechatAlertType,
		Enable:          &c.Enable,
		LastModified:    time.Now(),
	}
}
