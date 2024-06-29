package json

type Dict struct {
	Id       string  `json:"id"`
	Label    string  `json:"label"`
	Type     string  `json:"type"`
	Children []*Dict `json:"children" gorm:"-"`
}
