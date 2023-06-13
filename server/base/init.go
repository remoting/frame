package base

import (
	"errors"
	"fmt"

	"github.com/remoting/frame/json"
	"github.com/remoting/frame/server/web"
	"github.com/remoting/frame/util"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB(dbType, dbDsn string, models []any) (*gorm.DB, error) {
	var err error
	if dbType == "sqlite" {
		db, err = gorm.Open(sqlite.Open(dbDsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if dbType == "mysql" {
		db, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	} else {
		err = errors.New("dbType error")
	}
	if err != nil {
		fmt.Print("数据库连接错误:", err)
		return nil, err
	}
	for _, model := range models {
		err1 := db.AutoMigrate(model)
		if err != nil {
			fmt.Print("数据库初始化错误:", err1)
			return nil, err1
		}
	}
	return db, nil
}

type BaseModel struct {
	Id        string         `json:"id" gorm:"primaryKey;type:varchar(32)"`
	CreatedAt json.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy"`
	UpdatedAt json.Time      `json:"updatedAt"`
	UpdatedBy string         `json:"updatedBy"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	DeletedBy string         `json:"deletedBy"`
}

func (model *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	model.Id = util.NewUUID()
	return
}

type BaseForm struct {
	Context *web.Context `json:"-" gorm:"-"`
}

func (form *BaseForm) BindContext(ctx *web.Context) {
	form.Context = ctx
}

type Model interface {
}
type Form interface {
	BindContext(ctx *web.Context)
	NewForm() Form
	NewModel() Model
	GetCreateModel() Model
	GetUpdateModel() Model
	GetSearch(filter *SearchFilter) (string, []any, string, []any)
	GetById(id any) (string, []any)
}
type Service[T Form] interface {
	GetDB() *gorm.DB
	GetFormById(id any) (T, error)
	Update(form T) error
	PageSearch(filter *SearchFilter) (*SearchPaging[T], error)
	DelById(id any) error
	Create(form T) error
}

type Paging struct {
	Total   int `json:"total"`
	Current int `json:"current"`
	Size    int `json:"size"`
}
type SearchFilter struct {
	*Paging `json:"paging"`
	Filter  map[string]any `json:"filter"`
}

func (filter *SearchFilter) Offset() int {
	return (filter.Current - 1) * filter.Size
}

type SearchPaging[T any] struct {
	Paging *Paging `json:"paging"`
	List   []T     `json:"list"`
}
