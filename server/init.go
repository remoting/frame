package server

import (
	"github.com/remoting/frame/pkg/config"
	"github.com/remoting/frame/pkg/json"
	"github.com/remoting/frame/server/db"
)

func OnInit(models []any) error {
	_, err := db.InitDB(models)
	if err != nil {
		return err
	}
	err = InitConf()
	if err != nil {
		return err
	}
	return nil
}

func InitConf() error {
	list, err := db.Query[[]*json.Dict](` select conf_key as id,conf_val as label from sys_config  `)
	if err != nil {
		return err
	}
	for _, dict := range list {
		config.PutConfig(dict.Id, dict.Label)
	}
	return nil
}
