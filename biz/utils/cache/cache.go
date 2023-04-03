package cache

import (
	"encoding/json"
	"jx-hook/biz/config"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var client *redis.Client
var redisConfig config.RedisConfig

const (
	ALERT_PREFIX  = "ALERT_PREFIX:"
	SENDER_PREFIX = "SENDER_PREFIX:"
)

func init() {
	redisConfig = config.ConfigInstance.RedisConfig
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		hlog.Error("Error init redis client", err)
		os.Exit(1)
	}
}

// Create a normal cache method
func Set(key string, value any, expiration time.Duration) error {
	jsonVal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(key, jsonVal, expiration).Err()
}

// Get method to retrieve a value from cache
func Get(key string, res any) error {
	val, err := client.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), &res)
}

// Remove method to delete a key from cache
func Remove(key string) error {
	return client.Del(key).Err()
}

// Note: This assumes that the redis client has already been initialized in the init() function.
