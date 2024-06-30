package web

type Menu interface {
	GetId() string
	GetIcon() string
	GetLabel() string
	GetPrefix() string
	GetType() string
	GetRoute() string
	GetSubRoute() string
	GetChildren() []Menu
	GetParentId() string
}
type Role interface {
	GetId() string
	GetName() string
}
type Tenant interface {
	GetId() string
	GetName() string
	GetRoles() []Role
	IsOwner() bool
}
type User interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetMenus() []Menu
	GetTenants() []Tenant
	IsAdministrator() bool
}
