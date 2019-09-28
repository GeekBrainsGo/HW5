package main

import (
	"database/sql"
	"flag"
	"serv/server"

	_ "github.com/go-sql-driver/MySQL"
	"github.com/volatiletech/sqlboiler/boil"
)

const staticDir = "www/static"

func main() {
	flagServAddr := flag.String("addr", "localhost:8080", "server address")
	flagConnDb := flag.String("conndb", "mysql:mysql123!@tcp(192.168.99.100:3306)/blog", "db conn string")

	lg := NewLogger()
	db, err := sql.Open("mysql", *flagConnDb)
	boil.SetDB(db)
	if err != nil {
		lg.Panic("Can't connect to DB", err)
	} else {
		lg.Info("Connection to DB successful")
	}

	lg.Info(db.Ping())

	//blog := models.Blog{}

	srv := server.New(lg, db, staticDir)
	srv.Start(*flagServAddr)
}
