package models

import (
	"github.com/jinzhu/gorm"
)

// PostItem - объект Post
type PostItem struct {
	gorm.Model
	ID               uint   `gorm:"column:id"`
	Title            string `gorm:"column:title"`
	Dt               string `gorm:"column:dt"`
	SmallDescription string `gorm:"column:smalldescription"`
	Description      string `gorm:"column:description"`
}

// PostItemSlice - массив постов
type PostItemSlice []PostItem

// Insert - добавляет пост в БД
func (post *PostItem) Insert(db *gorm.DB) {

	postToInsert := PostItem{
		Title:            post.Title,
		Dt:               post.Dt,
		SmallDescription: post.SmallDescription,
		Description:      post.Description,
	}
	db.Create(&postToInsert)

}

// Delete - удалят объект из базы
func (post *PostItem) Delete(db *gorm.DB) {
	db.Exec(
		"DELETE FROM posts WHERE ID = ?",
		post.ID,
	)
}

// Update - изменяет пост в БД
func (post *PostItem) Update(db *gorm.DB) {

	db.Exec(
		"UPDATE post_items SET title = ?, dt = ?, smalldescription = ?, description = ? WHERE ID = ?",
		post.Title, post.Dt, post.SmallDescription, post.Description, post.ID,
	)

}

// GetAllPosts - получение всех постов
func GetAllPosts(db *gorm.DB) PostItemSlice {
	postsToFind := PostItemSlice{}
	db.Find(&postsToFind)
	return postsToFind
}
