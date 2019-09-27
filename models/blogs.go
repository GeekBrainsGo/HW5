package models

import (
	"github.com/jinzhu/gorm"
)

// BlogItem - объект блога
// type BlogItem struct {
// 	ID    int64  `json:"id"`
// 	Title string `json:"title"`
// 	Body  string `json:"article"`
// }

// BlogItemSlice - массив блогов
// type BlogItemSlice []BlogItem

// BlogItem - объект блога
type BlogItem struct {
	gorm.Model
	Title string `gorm:"column:title;type:varchar(45)"`
	Body  string `gorm:"column:body;default:''"`
}

type BlogItems []BlogItem

func (BlogItem) TableName() string {
	return "blogitems"
}
// AddBlog - обновляет объект в БД
func (blog *BlogItem) AddBlog(db *gorm.DB) error {

	db.Create(&blog)

	// _, err := db.Exec(
	// 	"INSERT INTO BlogItems (Title, Body) VALUES ( ?,  ? )",
	// 	blog.Title, blog.Body,
	// )
	// return err
	return nil
}

// UpdateBlog - обновляет объект в БД
func (blog *BlogItem) UpdateBlog(db *gorm.DB) error {

	db.Save(&blog)	
	// _, err := db.Exec(
	// 	"UPDATE BlogItems SET Title = ?, Body = ? WHERE ID = ?",
	// 	blog.Title, blog.Body, blog.ID,
	// )
	// return err
	return nil
}

// GetAllBlogItems - получение всех блогов
func GetAllBlogItems(db *gorm.DB) (BlogItems, error) {

	blogs := BlogItems{}
	db.Find(&blogs)

	// rows, err := db.Query("SELECT ID, Title, Body FROM BlogItems")
	// if err != nil {
	// 	return nil, err
	// }
	// blogs := make(BlogItems, 0, 8)
	// for rows.Next() {
	// 	blog := BlogItem{}
	// 	if err := rows.Scan(&blog.ID, &blog.Title, &blog.Body); err != nil {
	// 		return nil, err
	// 	}
	// 	blogs = append(blogs, blog)
	// }
	// return blogs, err
	return blogs,nil
}

// GetAllBlogItems - получение всех блогов
func GetBlogItem(db *gorm.DB, id uint) (BlogItem, error) {

	blog := BlogItem{}
	// var err error

	// row := db.QueryRow("SELECT ID, Title, Body FROM BlogItems WHERE ID = ?", id)

	// Get record with primary key (only works for integer primary key)
	db.First(&blog, blog.ID)
	//// SELECT * FROM blogitems WHERE id = NUM;
	// if row == nil {
	// 	return blog, errors.New("Пустое значение!")
	// }

	// if row != nil {
	// 	if err := row.Scan(&blog.ID, &blog.Title, &blog.Body); err != nil {
	// 		return blog, err
	// 	}
	// }
	// return blog, err
	return blog, nil
}

// Delete - удалят объект из базы
func (blog *BlogItem) Delete(db *gorm.DB) error {
	// _, err := db.Exec(
	// 	"DELETE FROM blogitems WHERE ID = ?",
	// 	blog.ID,
	// )
	// soft deleted
	db.Delete(&blog)	// UPDATE blogitems SET deleted_at="2013-10-29 10:23" WHERE id = NUM;
	// return err
	return nil
}

// Insert - добавляет блог в БД
func (blog *BlogItem) Insert(db *gorm.DB) error {
	// _, err := db.Exec(
	// 	"INSERT INTO BlogItems (ID, Title) VALUES (?, ?)",
	// 	blog.ID, blog.Title,
	// )
	db.Create(&blog)
	// return err
	return nil
}
