// Package ormblog implement mysql blog server with gorm orm.
package main

/*
	Basics Go.
	Rishat Ishbulatov, dated Sep 26, 2019.
	Write functions for creating, editing and deleting data in the database.
	Rewrite your service using BeeGo (create controllers and write algorithms
	for working with data through the website). Make a separate branch in Git
	and rewrite the work with the database using BeeGo orm. Transfer your blog
	to one of ORM: sqlboiler || gorm || beego-orm.
*/

import (
	"HW5/ormblog/server"
	"flag"
	"os"
	"os/signal"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	flagRootDir := flag.String("rootdir", "./www", "root dir of the server")
	flagServAddr := flag.String("addr", "localhost:8080", "server address")
	flag.Parse()

	lg := NewLogger()
	db, err := gorm.Open("mysql", "mysql:root@/ormblog?parseTime=true")
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}
	defer db.Close()

	serv := server.New(lg, *flagRootDir, db)
	go func() {
		err := serv.Start(*flagServAddr)
		if err != nil {
			lg.WithError(err).Fatal("can't run the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
}

// NewLogger creates new logger.
func NewLogger() *logrus.Logger {
	lg := logrus.New()
	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})
	lg.SetLevel(logrus.DebugLevel)
	return lg
}
