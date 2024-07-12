package config

import (
	"encoding/json"
	"github.com/remoting/frame/pkg/conv"
	"os"
	"strings"
	"sync"

	"github.com/remoting/frame/pkg/logger"
)

var (
	Value *_config
)

type Database struct {
	Type   string   `json:"type"`
	Master string   `json:"master"`
	Slave  []string `json:"slave"`
}
type _config struct {
	File     string         `json:"-"`
	Prefix   string         `json:"prefix,omitempty"`
	Database *Database      `json:"database,omitempty"`
	Version  string         `json:"version"`
	UiDir    string         `json:"ui-dir,omitempty"`
	Bind     string         `json:"bind"`
	Custom   map[string]any `json:"custom"`
}

func InitOnStart(file string) {
	Value = &_config{
		File: file,
	}
	configBytes, err := os.ReadFile(file)
	if err != nil {
		logger.Warn("file not found=%s", file)
	}
	err = json.Unmarshal(configBytes, &Value)
	if err != nil {
		logger.Error("error:%v", err)
	}
	if Value.Bind == "" {
		Value.Bind = "0.0.0.0:6383"
	}
}

/////
// 以下内容是保存数据库中的配置项
////

var lock = sync.RWMutex{}

func GetConfig(name string) string {
	lock.RLock()
	defer lock.RUnlock()
	return getConfig(name, Value.Custom)
}
func getConfig(name string, vars map[string]any) string {
	names := strings.Split(name, ".")
	if len(names) <= 1 {
		val, ok := vars[names[0]]
		if ok {
			return conv.String(val)
		}
	} else {
		val, ok := vars[names[0]]
		if ok {
			if ret, _ok := val.(map[string]interface{}); _ok {
				return getConfig(strings.Join(names[1:], "."), ret)
			} else {
				return ""
			}
		} else {
			logger.Warn("config format error")
			return ""
		}
	}
	return ""
}
func Save() error {
	lock.Lock()
	defer lock.Unlock()
	jsonStr, err := json.MarshalIndent(Value, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(Value.File, jsonStr, 0644)
}
func PutConfig(name, val string) {
	lock.Lock()
	defer lock.Unlock()
	if Value.Custom == nil {
		Value.Custom = make(map[string]any, 0)
	}
	setConfig(name, val, Value.Custom)
}
func setConfig(name, val string, vars map[string]any) {
	names := strings.Split(name, ".")
	if len(names) <= 1 {
		vars[name] = val
	} else {
		_vars, ok := vars[names[0]]
		if ok {
			if ret, _ok := _vars.(map[string]interface{}); _ok {
				setConfig(strings.Join(names[1:], "."), val, ret)
			} else {
				logger.Error("===config format error===")
			}
		} else {
			x_vars := make(map[string]any, 0)
			vars[names[0]] = x_vars
			setConfig(strings.Join(names[1:], "."), val, x_vars)
		}
	}
}
