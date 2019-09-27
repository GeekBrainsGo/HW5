package models

import (
	"github.com/jinzhu/gorm"
)

// BlogItem - объект блога
type BlogItem struct {
	gorm.Model
	Title string `gorm:"column:title;type:varchar(45)"`
	Body  string `json:"article" gorm:"column:body;default:''"`
}

// BlogItems - список блогов
type BlogItems []BlogItem

// TableName - имя таблицы блогов
func (BlogItem) TableName() string {
	return "blogitems"
}

// AddBlog - обновляет объект в БД
func (blog *BlogItem) AddBlog(db *gorm.DB) error {

	db.Create(&blog)

	return nil
}

// UpdateBlog - обновляет объект в БД
func (blog *BlogItem) UpdateBlog(db *gorm.DB) error {

	db.Save(&blog)

	return nil
}

// GetAllBlogItems - получение всех блогов
func GetAllBlogItems(db *gorm.DB) (BlogItems, error) {

	blogs := BlogItems{}
	db.Find(&blogs)

	return blogs, nil
}

// GetAllBlogItems - получение всех блогов
func GetBlogItem(db *gorm.DB, id uint) (BlogItem, error) {

	blog := BlogItem{}

	// Get record with primary key (only works for integer primary key)
	db.First(&blog, id)

	return blog, nil
}

// Delete - удалят объект из базы
func (blog *BlogItem) Delete(db *gorm.DB) error {

	// soft deleted
	db.Delete(&blog) // UPDATE blogitems SET deleted_at="2013-10-29 10:23" WHERE id = NUM;
	return nil
}
