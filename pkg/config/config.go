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

type Config struct {
	Prefix  string `json:"prefix"`
	Port    int    `json:"port"`
	Version string `json:"version"`
	DbType  string `json:"db-type"`
	DbFile  string `json:"db-file"`
	UiDir   string `json:"ui-dir"`
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
	if Value.Port == 0 {
		Value.Port = 8086
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
