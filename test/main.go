package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

// BlogItem - объект блога
type BlogItem struct {
	gorm.Model
	// ID    uint64 `gorm:"primary_key"`
	Title string `gorm:"column:title;type:varchar(45)"`
	Body  string `gorm:"column:body;default:''"`
}

type BlogItems []BlogItem

func (BlogItem) TableName() string {
	return "blogitems"
}

type DSN struct {
	Driver   string
	User     string
	Password string
	Protocol string
	Host     string
	DB       string
}

func (dsn DSN) Name() string {
	var ret string
	ret = dsn.User + ":" + dsn.Password + "@"
	if dsn.Protocol != "" {
		ret += dsn.Protocol + "("
	}
	ret += dsn.Host
	if dsn.Protocol != "" {
		ret += ")"
	}
	ret += "/" + dsn.DB
	return ret
}

func main() {

	dsn := DSN{
		Driver:   "mysql",
		User:     "root",
		Password: "root",
		Protocol: "tcp",
		Host:     "MASTER01",
		DB:       "MyBlogs",
	}

	fmt.Println(dsn.Name())

	db, err := gorm.Open("mysql", "root:root@tcp(MASTER01)/MyBlogs?parseTime=true")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connection is Ok!")
	}
	// Migrate the schema
	// db.AutoMigrate(&BlogItem{})

	blog :=BlogItem{}
	blogs :=BlogItems{}
	// usersToFind := models.Users{}

	db.First(&blog, 1) // find blog with id 1
	fmt.Println(blog)
	// db.First(&blog, "id = ?", 2) // find blog with code 2

	blog =BlogItem{}
	db.First(&blog, 2) // find blog with code 2
	// db.Find(&blog, 2)
	fmt.Println("ID второй записи:",blog.ID)
	// fmt.Println("ID:",blog.ID)
	fmt.Println(blog.Title)
	db.Find(&blogs)
	for _,item := range blogs{
		fmt.Println(item.ID, item.Title)
	}
	// fmt.Println(blogs)

	db.Model(&BlogItem{}).Where("id = ?", "1").Update("title", "My first Blog")
	blog =BlogItem{}
	db.First(&blog, 1) // find blog with id 1
	fmt.Println(blog)
}
