package model

import (
	"doovvvblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(20);not null" json:"title"`
	Category Category `gorm:"foreignkey:Cid"`
	Cid      int      `gorm:"type:int;not null;" json:"cid"`
	Desc     string   `gorm:"type:varchar(200);not null;" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

func (a *Article) TableName() string {
	return "article"
}
func CreateArt(a *Article) int {
	err = DB.Create(&a).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询所有文章
func GetArts(pageSize int, pageNum int) ([]Article, int64) {
	var arts []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := DB.Preload("Category").Limit(pageSize).Offset(offset).Find(&arts).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return arts, total
}

// 查询单个文章
func GetArt(id int) (Article, int) {
	var art Article
	err := DB.Preload("Category").Where("id =?", id).First(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// 查询分类下的所有文章
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var arts []Article
	var total int64

	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid =?", cid).Find(&arts).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return arts, errmsg.SUCCESS, total

}

// 删除文章
func DeleteArt(id int) (code int) {
	var art Article
	err := DB.Where("id =?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户
func EditArt(id uint, a *Article) (code int) {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = a.Title
	maps["cid"] = a.Cid
	maps["desc"] = a.Desc
	maps["content"] = a.Content
	maps["img"] = a.Img

	err := DB.Model(&art).Where("id =?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
