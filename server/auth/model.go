package auth

import "github.com/remoting/frame/server/web"

type UserImpl struct {
	Id      string
	Name    string
	Email   string
	Menus   []web.Menu
	Tenants []web.Tenant
}

func (user *UserImpl) GetId() string {
	return user.Id
}
func (user *UserImpl) GetName() string {
	return user.Name
}
func (user *UserImpl) GetEmail() string {
	return user.Email
}
func (user *UserImpl) GetMenus() []web.Menu {
	return user.Menus
}
func (user *UserImpl) GetTenants() []web.Tenant {
	return user.Tenants
}
func (user *UserImpl) IsAdministrator() bool {
	return true
}

type TenantImpl struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Roles []web.Role
}

func (tenant *TenantImpl) GetId() string {
	return tenant.Id
}
func (tenant *TenantImpl) GetName() string {
	return tenant.Name
}
func (tenant *TenantImpl) GetRoles() []web.Role {
	return tenant.Roles
}
func (tenant *TenantImpl) IsOwner() bool {
	for _, role := range tenant.Roles {
		if role.GetId() == "owner" {
			return true
		}
	}
	return false
}

type RoleImpl struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (role *RoleImpl) GetId() string {
	return role.Id
}
func (role *RoleImpl) GetName() string {
	return role.Name
}

type MenuImpl struct {
	Id       string     `json:"id"`
	Icon     string     `json:"icon"`
	Label    string     `json:"label"`
	Prefix   string     `json:"prefix"`
	Type     string     `json:"type"`
	Route    string     `json:"route"`
	SubRoute string     `json:"subRoute"`
	Children []web.Menu `json:"children"`
	ParentId string     `json:"parentId"`
}

func (menu *MenuImpl) GetId() string {
	return menu.Id
}
func (menu *MenuImpl) GetLabel() string {
	return menu.Label
}
func (menu *MenuImpl) GetIcon() string {
	return menu.Icon
}
func (menu *MenuImpl) GetPrefix() string {
	return menu.Prefix
}
func (menu *MenuImpl) GetType() string {
	return menu.Type
}
func (menu *MenuImpl) GetRoute() string {
	return menu.Route
}
func (menu *MenuImpl) GetSubRoute() string {
	return menu.SubRoute
}
func (menu *MenuImpl) GetChildren() []web.Menu {
	return menu.Children
}
func (menu *MenuImpl) GetParentId() string {
	return menu.ParentId
}
