package model

import (
	"doovvvblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null;" json:"name"`
}

func (Category) TableName() string {
	return "category"
}
func CheckCategory(name string) (code int) {
	var cate Category
	err := DB.Select("id").Where("name = ?", name).Limit(1).Find(&cate).Error
	if err == gorm.ErrRecordNotFound {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}
func CreateCate(c *Category) int {
	err = DB.Create(&c).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func GetCates(pageSize int, pageNum int) []Category {
	var cates []Category
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := DB.Limit(pageSize).Offset(offset).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}
func DeleteCate(id int) (code int) {
	var cate Category
	err := DB.Where("id =?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户
func EditCate(id uint, c *Category) (code int) {
	var cate Category
	var maps = make(map[string]any)
	maps["name"] = c.Name
	maps["id"] = c.Id
	err := DB.Model(&cate).Where("id =?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
