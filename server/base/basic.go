package base

import (
	"github.com/remoting/frame/server/web"
	"gorm.io/gorm"
)

type IModel interface {
	NewModel() IModel
}
type IForm interface {
	NewForm() IForm
	GetModel() IModel
	BindContext(ctx *web.Context)
}
type IBasicService[T IModel] interface {
	GetDB() *gorm.DB
	GetFormById(id any) (T, error)
	Update(model T) error
	DelById(id any) error
	Create(model T) error
}
type BasicForm struct {
	Context *web.Context `json:"-" gorm:"-"`
}

func (form *BasicForm) BindContext(ctx *web.Context) {
	form.Context = ctx
}

// BasicController controller
type BasicController[T IModel, F IForm] struct {
	BasicService IBasicService[T]
}

func (controller *BasicController[T, F]) Bind(c *web.Context) F {
	var form F
	form = form.NewForm().(F)
	form.BindContext(c)
	c.BindOBJ(form)
	return form
}

func (controller *BasicController[T, F]) Add(c *web.Context) {
	form := controller.Bind(c)
	model := form.GetModel()
	err := controller.BasicService.Create(model.(T))
	c.CheckError(err)
	c.Success("")
}

func (controller *BasicController[T, F]) Update(c *web.Context) {
	form := controller.Bind(c)
	model := form.GetModel()
	err := controller.BasicService.Update(model.(T))
	c.CheckError(err)
	c.Success("")
}

func (controller *BasicController[T, F]) Delete(c *web.Context) {
	json := c.ParseJSON()
	err := controller.BasicService.DelById(json.GetString("id"))
	c.CheckError(err)
	c.Success("")
}

func (controller *BasicController[T, F]) View(c *web.Context) {
	json := c.ParseJSON()
	form, err := controller.BasicService.GetFormById(json.GetString("id"))
	c.CheckError(err)
	c.Success(form)
}

// BasicService service
type BasicService[T IModel] struct {
}

func (service *BasicService[T]) GetDB() *gorm.DB {
	return GetDB()
}

func (service *BasicService[T]) DelById(id any) error {
	var model T
	return service.GetDB().Where("id=?", id).Delete(model.NewModel()).Error
}

func (service *BasicService[T]) GetFormById(id any) (T, error) {
	var model T
	ret := model.NewModel()
	if err := service.GetDB().First(ret, "id = ?", id).Error; err != nil {
		return model, err
	}
	return ret.(T), nil
}

func (service *BasicService[T]) Create(model T) error {
	result := service.GetDB().Create(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *BasicService[T]) Update(model T) error {
	result := service.GetDB().Updates(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
