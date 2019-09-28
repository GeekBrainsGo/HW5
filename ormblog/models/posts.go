package models

import (
	"html/template"

	"github.com/jinzhu/gorm"
)

// Post stands for post object.
type Post struct {
	gorm.Model
	Title   string        `gorm:"column:title;type:varchar(16)"`
	Author  string        `gorm:"column:author;type:varchar(16)"`
	Content template.HTML `gorm:"column:content;type:text"`
}

// Posts stands for array of posts.
type Posts []Post

// Insert post to database.
func (p *Post) Insert(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// Delete deletes post from database.
func (p *Post) Delete(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Delete(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// Update updates post in database.
func (p *Post) Update(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Update(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// AllPosts return all posts from database.
func AllPosts(db *gorm.DB) (Posts, error) {
	posts := Posts{}
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
