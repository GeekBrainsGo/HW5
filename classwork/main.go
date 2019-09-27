package main

import (
	"fmt"
	"log"
	"orm/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// db, err := gorm.Open("mysql", "mysql:root@/beego?parseTime=true")
	db, err := gorm.Open("mysql", "root:root@tcp(MASTER01)/myuser?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()

	// Создать пользователя

	userToCreate := models.User{
		Name:     "Andrew",
		LastName: "Ivanov",
		Access:   "admin",
	}

	db.Create(&userToCreate)

	// Найти пользователя по уровню доступа

	userToFind := models.User{}
	db.Find(&userToFind, "access = ?", "admin")
	fmt.Println("userToFind", userToFind)

	// Найти пользователей по уровню доступа

	usersToFind := models.Users{}
	db.Find(&usersToFind, "access = ?", "admin")
	// fmt.Println("usersToFind", usersToFind)
	for _,item :=range usersToFind {
		fmt.Println("user:", item.ID,item.LastName)
	}

	// Обновить уровень доступа всем пользователям у которых его нет

	// db.Model(&models.User{}).Where("access = ?", "").Update("access", "guest")
}
