package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Post - объект поста в блоге
type Post struct {
	gorm.Model
	ID        uint `gorm:"primary_key;column:id" json:"id"`
	Title     string `gorm:"column:title" json:"title"`
	Text      string `gorm:"column:text" json:"text"`
	Labels    []string `json:"labels"`
}

// Posts - массив постов в блоге
type Posts []Post

// Create - создает задачу в БД
func (post *Post) Create(db *gorm.DB) (Post, error) {

	db.Create(&post)
	if db.Error != nil {
		return Post{}, db.Error
	}

	var newPost Post
	db.First(&newPost, post.ID)

	return newPost, db.Error
}

// Delete - удалить объект из базы
func (post *Post) Delete(db *gorm.DB) error {
	db.Delete(&post)

	return db.Error
}

// Update - обновляет объект в БД
func (post *Post) Update(db *gorm.DB) error {
	db.Model(&post).Updates(Post{Title: post.Title, Text: post.Text})

	return db.Error
}

// GetPost - получение поста
func GetPost(postId uint, db *gorm.DB) (Post, error) {

	var getPost Post
	db.First(&getPost, postId)

	fmt.Println(getPost.ID)
	fmt.Println(getPost.Title)

	return getPost, db.Error
}

// GetAllPosts - получение всех постов
func GetAllPosts(db *gorm.DB) (Posts, error) {

	posts := Posts{}
	db.Find(&posts)

	return posts, db.Error
}
