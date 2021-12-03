package models

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

// article
type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func CreateArticle(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个文章
func GetArtInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

//查询分类下所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var article []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&article).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATENAME_NOT_EXIST, 0
	}
	return article, errmsg.SUCCESS, total
}

//查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int, int) {
	var article []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return article, errmsg.SUCCESS, total
}

//编辑文章
func EditArticle(id int, article *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img
	err := db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
