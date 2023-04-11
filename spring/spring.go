package spring

import (
	"reflect"
)

type Spring interface {
	Reg(bean any)
	Get(name string) any
	New() Spring
}
type context struct {
	parent Spring
	beans  map[string]any
}

func New(parent Spring) Spring {
	c := &context{
		beans: make(map[string]any),
	}
	if parent != nil {
		c.parent = parent
	}
	return c
}
func (ctx *context) Reg(bean any) {
	t := reflect.TypeOf(bean)
	name := t.Elem().PkgPath() + "/" + t.Elem().Name()
	ctx.beans[name] = bean
}
func (ctx *context) New() Spring {
	c := &context{
		beans: make(map[string]any),
	}
	c.parent = ctx
	return c
}
func (ctx *context) Get(name string) any {
	bean, ok := ctx.beans[name]
	if ok {
		return bean
	} else {
		if ctx.parent != nil {
			return ctx.parent.Get(name)
		} else {
			panic("bean not found: [ " + name + " ]")
		}
	}
}
func GetBeanName[T any]() string {
	var t T
	t1 := reflect.TypeOf(t)
	return t1.Elem().PkgPath() + "/" + t1.Elem().Name()
}
func GetBean[T any](spring Spring) T {
	return spring.Get(GetBeanName[T]()).(T)
}
