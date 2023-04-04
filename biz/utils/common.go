package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jx-hook/biz/config"
	"jx-hook/biz/models/senderConfig"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var dictMaxSize = 2000
var client *redis.Client

func UpdateDictMaxSize(size int) {
	hlog.Info("Update dict max size ", size)
	dictMaxSize = size
}

func UpdateClient(redisConfig config.RedisConfig) {
	hlog.Info("Update redis config ", redisConfig)
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})
	err := client.Ping().Err()
	if err != nil {
		hlog.Error("Can't access redis server with 'ping' cmd")
	}
}

// cache prefix
const (
	AlertPrefix  = "ALERT_PREFIX:"
	SenderPrefix = "SENDER_PREFIX:"
)

// MSG TYPE
const (
	WechatMarkdown = "wechat.markdown"
	WechatText     = "wechat.text"
	Wechat         = "wechat"
	Custom         = "custom"
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

// Send msg to url according to sender config
func SendMsg(config senderConfig.SenderConfig, msg string) error {
	msgType, robotKey, customUrl := config.AlertType, config.WechatRobotKey, config.CustomUrl

	switch msgType {
	case WechatText:
		return SendTextMsg(robotKey, msg)
	case WechatMarkdown:
		return SendTextMsg(robotKey, msg)
	default:
		sendByte := []byte(msg)
		if strings.HasPrefix(msgType, Wechat) {
			return sendWechat(robotKey, sendByte)
		} else {
			return sendHttp(customUrl, sendByte)
		}
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
	return sendWechat(robotKey, msgJSON)
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
	return sendWechat(robotKey, msgJSON)
}

func sendWechat(robotKey string, msg []byte) error {
	return sendHttp(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", robotKey), msg)
}

func sendHttp(url string, msg []byte) error {
	_, err := http.Post(url, "application/json", bytes.NewReader(msg))
	if err != nil {
		hlog.Error("failed send msg to ", url, err)
		return err
	}
	hlog.Info("Succeed send msg to ", url)
	return nil
}

func ResolveTemplate(t string, jsonByte []byte) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonByte, &data)
	if err != nil {
		hlog.Error("Failed to deserialize data")
		return "", err
	}

	trySize := 0

	var innerResolve func(string, map[string]interface{})
	innerResolve = func(parent string, d map[string]interface{}) {
		for k, val := range d {
			// in case maxSize cycle
			trySize++
			if trySize > dictMaxSize {
				return
			}

			// get the current key
			var key string
			if parent == "" {
				key = k
			} else {
				key = parent + "." + k
			}

			// parse value's type
			switch v := val.(type) {
			case map[string]interface{}:
				hlog.Debug("Doing key ", key, " as interface resolve & type :", v)
				innerResolve(key, val.(map[string]interface{}))
			default:
				placeholder := "${" + key + "}"
				t = strings.Replace(t, placeholder, fmt.Sprint(val), -1)
			}
		}
	}

	innerResolve("", data)
	return t, nil
}

// Create a normal cache method
func Cache(key string, value any, expiration time.Duration) error {
	jsonVal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(key, jsonVal, expiration).Err()
}

// Get method to retrieve a value from cache
func GetCache(key string, res any) error {
	val, err := client.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), &res)
}

// Remove method to delete a key from cache
func RemoveCache(key string) error {
	return client.Del(key).Err()
}
