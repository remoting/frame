package base

import (
	"gorm.io/gorm"
)

type BaseService[T Form] struct {
}

func (service *BaseService[T]) GetDB() *gorm.DB {
	return GetDB()
}
func GetDB() *gorm.DB {
	return db
}
func (service *BaseService[T]) PageSearch(filter *SearchFilter) (*SearchPaging[T], error) {
	var list []T
	var total int
	var t T
	a, b, c, d := t.GetSearch(filter)
	if err := service.GetDB().Raw(a, b...).Scan(&total).Error; err != nil {
		return nil, err
	}
	if err := service.GetDB().Raw(c, d...).Scan(&list).Error; err != nil {
		return nil, err
	}
	return &SearchPaging[T]{
		Paging: &Paging{
			Total:   total,
			Current: filter.Current,
			Size:    filter.Size,
		},
		List: list,
	}, nil
}

func (service *BaseService[T]) DelById(id any) error {
	var form T
	return service.GetDB().Where("id=?", id).Delete(form.NewModel()).Error
}

func (service *BaseService[T]) GetFormById(id any) (T, error) {
	var form T
	a, b := form.GetById(id)
	if err := service.GetDB().Raw(a, b...).First(&form).Error; err != nil {
		return form, err
	}
	return form, nil
}

func (service *BaseService[T]) Create(form T) error {
	result := service.GetDB().Create(form.GetCreateModel())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *BaseService[T]) Update(form T) error {
	model := form.GetUpdateModel()
	result := service.GetDB().Updates(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
