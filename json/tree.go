package json

type Node struct {
	Id       string  `json:"id"`
	Label    string  `json:"label"`
	Type     string  `json:"type"`
	Route    string  `json:"route"`
	SubRoute string  `json:"subRoute"`
	Icon     string  `json:"icon"`
	ParentId string  `json:"parentId"`
	Children []*Node `json:"children" gorm:"-"`
	Show     bool    `json:"-"`
}
