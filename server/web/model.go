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
type Tenant interface {
	GetId() string
	GetName() string
	GetRoles() []string
	IsOwner() bool
	IsAdmin() bool
}
type User interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetMenus() []Menu
	GetTenants() []Tenant
	GetTenantId() string
	IsAdministrator() bool
}
