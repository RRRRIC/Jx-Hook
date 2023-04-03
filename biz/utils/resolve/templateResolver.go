package resolve

import (
	"encoding/json"
	"fmt"
	"jx-hook/biz/config"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var dictMaxSize int

func init() {
	dictMaxSize = config.ConfigInstance.DictMaxSize
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
			switch v := val.(type) {
			case map[string]interface{}:
				hlog.Info("Doing key ", key, " as interface resolve & type :", v)
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
