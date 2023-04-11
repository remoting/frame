package base

import (
	"github.com/remoting/frame/server/web"
)

type BaseController[T Form] struct {
	Service Service[T]
}

func (controller *BaseController[T]) Bind(c *web.Context) T {
	var form T
	form = form.NewForm().(T)
	form.BindContext(c)
	c.BindOBJ(form)
	return form
}

func (controller *BaseController[T]) Add(c *web.Context) {
	form := controller.Bind(c)
	err := controller.Service.Create(form)
	c.CheckError(err)
	c.Success("")
}

func (controller *BaseController[T]) Update(c *web.Context) {
	form := controller.Bind(c)
	err := controller.Service.Update(form)
	c.CheckError(err)
	c.Success("")
}
func (controller *BaseController[T]) List(c *web.Context) {
	filter := &SearchFilter{}
	c.BindOBJ(filter)
	rows, err := controller.Service.PageSearch(filter)
	c.CheckError(err)
	c.Success(rows)
}

func (controller *BaseController[T]) Delete(c *web.Context) {
	json := c.ParseJSON()
	err := controller.Service.DelById(json.GetString("id"))
	c.CheckError(err)
	c.Success("")
}

func (controller *BaseController[T]) View(c *web.Context) {
	json := c.ParseJSON()
	form, err := controller.Service.GetFormById(json.GetString("id"))
	c.CheckError(err)
	c.Success(form)
}
