package auth

import "github.com/remoting/frame/server/web"

// UserImpl 用户结构体
type UserImpl struct {
	Id            string       `json:"id"`
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Phone         string       `json:"phone"`
	Menus         []web.Menu   `json:"menus"`
	Tenants       []web.Tenant `json:"tenants"`
	Administrator bool         `json:"admin"`
	TenantId      string       `json:"tenantId"`
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
	return user.Administrator
}
func (user *UserImpl) GetTenantId() string {
	return user.TenantId
}

// TenantImpl 租户结构体
type TenantImpl struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Logo  string   `json:"logo"`
	Roles []string `json:"roles"`
}

func (tenant *TenantImpl) GetId() string {
	return tenant.Id
}
func (tenant *TenantImpl) GetName() string {
	return tenant.Name
}
func (tenant *TenantImpl) GetRoles() []string {
	return tenant.Roles
}
func (tenant *TenantImpl) IsOwner() bool {
	for _, role := range tenant.Roles {
		if role == "owner" {
			return true
		}
	}
	return false
}
func (tenant *TenantImpl) IsAdmin() bool {
	for _, role := range tenant.Roles {
		if role == "admin" {
			return true
		}
	}
	return false
}

// MenuImpl 菜单结构体
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
