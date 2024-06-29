package web

type Controller interface {
	OnInit(HookFunc)
}

type HookFunc func(string) *RouterGroup
