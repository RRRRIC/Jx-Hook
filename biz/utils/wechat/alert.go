package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// MSG TYPE
const (
	MARKDOWN = "markdown"
	TEXT     = "text"
	OTHER    = "OTHER"
)

// Define a struct to hold the message data
type message struct {
	MsgType  string   `json:"msgtype"`
	Text     text     `json:"text"`
	Markdown markdown `json:"markdown"`
}

// Define a struct to hold the text data
type text struct {
	Content string `json:"content"`
}

// Define a struct to hold the markdown data
type markdown struct {
	Content string `json:"content"`
}

// Send msg to wechat robot according to msg Type
func SendMsg(robotKey, msgType, msg string) error {
	switch msgType {
	case TEXT:
		return SendTextMsg(robotKey, msg)
	case MARKDOWN:
		return SendTextMsg(robotKey, msg)
	default:
		return SendDefault(robotKey, msg)
	}
}

// Define a function to send a message of type "msg"
func SendTextMsg(robotKey string, msg string) error {
	// Create the message data
	msgData := message{
		MsgType: "text",
		Text: text{
			Content: msg,
		},
	}
	// Convert the message data to JSON
	msgJSON, err := json.Marshal(msgData)
	if err != nil {
		return err
	}
	return send(robotKey, msgJSON)
}

// Define a function to send a message of type "markdown"
func SendMarkdown(robotKey string, msg string) error {
	// Create the message data
	msgData := message{
		MsgType: "markdown",
		Markdown: markdown{
			Content: msg,
		},
	}

	// Convert the message data to JSON
	msgJSON, err := json.Marshal(msgData)
	if err != nil {
		return err
	}
	return send(robotKey, msgJSON)
}

// Define a function to send a message of type "default"
func SendDefault(robotKey string, msg string) error {
	return send(robotKey, []byte(msg))
}

func send(robotKey string, msg []byte) error {
	_, err := http.Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", robotKey), "application/json", bytes.NewReader(msg))
	if err != nil {
		return err
	}
	return nil
}
