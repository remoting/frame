package db

import (
	"database/sql"
	"github.com/remoting/frame/pkg/config"
	"github.com/remoting/frame/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var __db *gorm.DB

func InitDB(models []any) (*gorm.DB, error) {
	var err error
	if config.Value.Database.Type == "sqlite" {
		__db, err = gorm.Open(sqlite.Open(config.Value.Database.Master), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if config.Value.Database.Type == "mysql" {
		__db, err = gorm.Open(mysql.Open(config.Value.Database.Master), &gorm.Config{})
		if err == nil {
			slave := make([]gorm.Dialector, 0)
			for i := 0; i < len(config.Value.Database.Slave); i++ {
				slave = append(slave, mysql.Open(config.Value.Database.Slave[i]))
			}
			if len(slave) > 0 {
				__db.Use(dbresolver.Register(dbresolver.Config{
					Replicas: slave,
				}))
			}
		}
	} else {
		err = errors.New("dbType error")
	}
	if err != nil {
		return nil, errors.Wrap(err, "数据库链接错误")
	}
	for _, model := range models {
		err1 := __db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(model)
		if err1 != nil {
			errors.Wrap(err1, "数据库初始化错误")
		}
	}
	return __db, nil
}

func Get() *Connection {
	return &Connection{
		DB: __db,
	}
}

type Connection struct {
	*gorm.DB
}

func (db *Connection) Clauses(conds ...clause.Expression) *Connection {
	return &Connection{
		DB: db.DB.Clauses(conds...),
	}
}
func (db *Connection) UseMaster() *Connection {
	return db.Clauses(dbresolver.Write)
}
func (db *Connection) UseSlave() *Connection {
	return db.Clauses(dbresolver.Read)
}
func (db *Connection) GetDB() *Connection {
	return Get()
}
func (db *Connection) Insert(table string, row map[string]interface{}) error {
	return db.DB.Table(table).Create(row).Error
}
func (db *Connection) Update(table string, row map[string]interface{}, where string, params ...interface{}) error {
	return db.DB.Table(table).Where(where, params...).Updates(row).Error
}
func (db *Connection) Execute(sql string, params ...interface{}) error {
	return db.DB.Exec(sql, params...).Error
}
func (db *Connection) Transaction(fc func(*Connection) error, opts ...*sql.TxOptions) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&Connection{
			DB: tx,
		})
	}, opts...)
}
func (db *Connection) Query(obj any, sql string, params ...interface{}) error {
	return db.DB.Raw(sql, params...).Find(obj).Error
}
func UseMaster() *Connection {
	return Get().UseMaster()
}
func UseSlave() *Connection {
	return Get().UseSlave()
}

// QueryFind 查询数据 如果未找到不会返回 Error
func QueryFind[T any](sql string, params ...interface{}) (T, error) {
	var obj T
	if err := Get().Raw(sql, params...).Find(&obj).Error; err != nil {
		return obj, err
	}
	return obj, nil
}

// QueryFirst 查询数据 如果未找到返回 Error
func QueryFirst[T any](sql string, params ...interface{}) (T, error) {
	var obj T
	if err := Get().Raw(sql, params...).First(&obj).Error; err != nil {
		return obj, err
	}
	return obj, nil
}
func Query[T any](sql string, params ...interface{}) (T, error) {
	var obj T
	if err := Get().Query(&obj, sql, params...); err != nil {
		return obj, err
	}
	return obj, nil
}
func ExecuteSQL(sql string, params ...interface{}) error {
	return Get().Execute(sql, params...)
}
func InsertRow(table string, row map[string]interface{}) error {
	return Get().Insert(table, row)
}
func UpdateRow(table string, row map[string]interface{}, where string, params ...interface{}) error {
	return Get().Update(table, row, where, params...)
}
func Transaction(fc func(tx *Connection) error, opts ...*sql.TxOptions) error {
	return Get().Transaction(fc, opts...)
}
