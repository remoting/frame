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
}
type User interface {
	GetId() string
	GetName() string
	GetMenus() []Menu
	GetRoles() []Role
	GetTenants() []Tenant
	IsAdministrator() bool
	IsTenantAdmin() bool
}
