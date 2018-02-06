package controller 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	db = getDb()
}

func getDb() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:joy139139@tcp(127.0.0.1:3306)/Ngdule?charset=utf8")
		checkErr(err)
		db.SetMaxIdleConns(2)
		db.SetMaxOpenConns(2)

		if err := db.Ping(); err != nil {
			log.Fatalln(err)
		}
	}
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
