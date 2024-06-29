package auth

import "github.com/remoting/frame/server/web"

type UserImpl struct {
	Id      string
	Name    string
	Menus   []web.Menu
	Roles   []web.Role
	Tenants []web.Tenant
}
type TenantImpl struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
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

func (user *UserImpl) GetId() string {
	return user.Id
}
func (user *UserImpl) IsAdministrator() bool {
	if user.Id == "administrator" {
		return true
	} else {
		for _, role := range user.Roles {
			if role.GetId() == "administrator" {
				return true
			}
		}
	}
	return false
}
func (user *UserImpl) IsTenantAdmin() bool {
	return true
}
