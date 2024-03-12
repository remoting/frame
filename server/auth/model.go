package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/json"
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
