package config

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/remoting/frame/pkg/logger"
)

var (
	Value Config
)

type Database struct {
	Type   string   `json:"type"`
	Master string   `json:"master"`
	Slave  []string `json:"slave"`
}
type Config struct {
	Prefix   string   `json:"prefix"`
	Database Database `json:"database"`
	Version  string   `json:"version"`
	UiDir    string   `json:"ui-dir"`
	Bind     string   `json:"bind"`
}

func InitOnStart(file string) {
	Value = Config{}
	configBytes, err := os.ReadFile(file)
	if err != nil {
		logger.Warn("file not found=%s", file)
	}
	err = json.Unmarshal(configBytes, &Value)
	if err != nil {
		logger.Error("error:%v", err)
	}
	if Value.Prefix == "" {
		Value.Prefix = "/"
	}
	if Value.Bind == "" {
		Value.Bind = "0.0.0.0:6383"
	}
}

/////
// 以下内容是保存数据库中的配置项
////

var setting = make(map[string]string)
var lock = sync.RWMutex{}

func GetConfig(name string) string {
	lock.RLock()
	defer lock.RUnlock()
	val, ok := setting[name]
	if ok {
		return val
	}
	return ""
}
func PutConfig(name, val string) {
	lock.Lock()
	defer lock.Unlock()
	setting[name] = val
}
