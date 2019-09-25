package models

import "github.com/jinzhu/gorm"

// User - структура пользователя
type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	LastName string `gorm:"column:last_name"`
	Access   string `gorm:"column:access"`
}

// Users - слайс пользователей
type Users []User
