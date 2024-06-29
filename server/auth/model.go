package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/json"
)

type Menu struct {
	Id       string  `json:"id"`
	Prefix   string  `json:"prefix"`
	Name     string  `json:"label"`
	Type     string  `json:"type"`
	Route    string  `json:"route"`
	SubRoute string  `json:"subRoute" gorm:"column:sub_route"`
	Icon     string  `json:"icon"`
	ParentId string  `json:"parentId" gorm:"column:parent_id"`
	Children []*Menu `json:"children" gorm:"-"`
	Show     bool    `json:"-"`
}

type UserInfo interface {
	UserId() string
	UserName() string
	GetMenus() []*Menu
	GetRoles() []*json.Dict
	GetTenant() []UserTenant
	IsAdmin() bool
}
type UserTenant interface {
	TenantId() string
	TenantName() string
}
type UserInfoService interface {
	GetUserInfoById(userId string) (UserInfo, error)
}

// AuthService  授权服务
type AuthService interface {
	Authorization(c *gin.Context) bool
}

// Authentication 认证服务
type Authentication struct {
	AuthService    AuthService     //授权服务
	UserService    UserInfoService //用户信息服务
	TokenSecretPub string          //jwt Token 验证公钥
	AnonymousPath  []string        //匿名可访问路径
}

type UserInfoImpl struct {
	Id     string            `json:"id"`
	Name   string            `json:"name"`
	Tenant []*UserTenantImpl `gorm:"-" json:"tenant"`
	Menus  []*Menu           `gorm:"-" json:"menus"`
	Roles  []*json.Dict      `gorm:"-" json:"roles"`
}
type UserTenantImpl struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func (tenant UserTenantImpl) TenantName() string {
	return tenant.Name
}
func (tenant UserTenantImpl) TenantId() string {
	return tenant.Id
}
func (user *UserInfoImpl) UserId() string {
	return user.Id
}
func (user *UserInfoImpl) UserName() string {
	return user.Name
}
func (user *UserInfoImpl) GetMenus() []*Menu {
	return user.Menus
}
func (user *UserInfoImpl) GetRoles() []*json.Dict {
	return user.Roles
}
func (user *UserInfoImpl) GetTenant() []UserTenant {
	var tenant []UserTenant
	for _, t := range user.Tenant {
		tenant = append(tenant, t)
	}
	return tenant
}
func (user *UserInfoImpl) IsAdmin() bool {
	if user.Id == "admin" {
		return true
	} else {
		for _, role := range user.Roles {
			if role.Id == "administrator" {
				return true
			}
		}
	}
	return false
}
