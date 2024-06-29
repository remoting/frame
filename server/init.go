package server

import (
	"github.com/remoting/frame/pkg/config"
	"github.com/remoting/frame/pkg/errors"
	"github.com/remoting/frame/pkg/json"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func OnInit(models []any) error {
	_, err := InitDB(models)
	if err != nil {
		return err
	}
	err = InitConf()
	if err != nil {
		return err
	}
	return nil
}
func InitDB(models []any) (*gorm.DB, error) {
	var err error
	if config.Value.DbType == "sqlite" {
		db, err = gorm.Open(sqlite.Open(config.Value.DbFile), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if config.Value.DbType == "mysql" {
		db, err = gorm.Open(mysql.Open(config.Value.DbFile), &gorm.Config{})
	} else {
		err = errors.New("dbType error")
	}
	if err != nil {
		return nil, errors.Wrap(err, "数据库链接错误")
	}
	for _, model := range models {
		err1 := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(model)
		if err1 != nil {
			errors.Wrap(err1, "数据库初始化错误")
		}
	}
	return db, nil
}

func InitConf() error {
	list, err := QueryList[*json.Dict](` select conf_key as id,conf_val as label from sys_config  `)
	if err != nil {
		return err
	}
	for _, dict := range list {
		config.PutConfig(dict.Id, dict.Label)
	}
	return nil
}
func GetDB() *gorm.DB {
	return db
}

// QueryFind 查询数据 如果未找到不会返回 Error
func QueryFind[T any](sql string, params ...interface{}) (T, error) {
	var obj T
	if err := GetDB().Raw(sql, params...).Find(&obj).Error; err != nil {
		return obj, err
	}
	return obj, nil
}

// QueryFirst 查询数据 如果未找到返回 Error
func QueryFirst[T any](sql string, params ...interface{}) (T, error) {
	var obj T
	if err := GetDB().Raw(sql, params...).First(&obj).Error; err != nil {
		return obj, err
	}
	return obj, nil
}
func QueryList[T any](sql string, params ...interface{}) ([]T, error) {
	var list []T
	if err := GetDB().Raw(sql, params...).Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func ExecuteSQL(sql string, params ...interface{}) error {
	if err := GetDB().Exec(sql, params...).Error; err != nil {
		return err
	}
	return nil
}
func InsertRow(table string, row map[string]interface{}) error {
	if err := GetDB().Table(table).Create(row).Error; err != nil {
		return err
	}
	return nil
}
func UpdateRow(table string, row map[string]interface{}, where string, params ...interface{}) error {
	if err := GetDB().Table(table).Where(where, params...).Updates(row).Error; err != nil {
		return err
	}
	return nil
}
