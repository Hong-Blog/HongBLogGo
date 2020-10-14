package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

func init() {
	//db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/oneblog")
	db, err := sqlx.Connect("mysql", "root:123qwe@tcp(127.0.0.1:3306)/oneblog?parseTime=true")
	if err != nil {
		log.Panicln("db err: ", err.Error())
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	Db = db
}
