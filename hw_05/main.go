package main

import (
	"flag"
	"os"
	"os/signal"
	"serv/logger"
	"serv/server"

	_ "github.com/go-sql-driver/MySQL"
	"github.com/jinzhu/gorm"
)

func main() {
	flagRootDir := flag.String("rootdir", "./www", "root dir of the server")
	flagServAddr := flag.String("addr", "localhost:8080", "server address")
	flag.Parse()

	lg := logger.NewLogger()

	db, err := gorm.Open("mysql", "mysql:root@/blog5?parseTime=true")
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}
	db = db.Debug()
	defer db.Close()

	serv := server.New(lg, *flagRootDir, db)

	go func() {
		err := serv.Start(*flagServAddr)
		if err != nil {
			lg.WithError(err).Fatal("can't run the server")
		}
	}()

	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, os.Interrupt, os.Kill)
	<-stopSig

}
