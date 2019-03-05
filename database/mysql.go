package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:smm123@tcp(127.0.0.1:3306)/world?parseTime=true")
	if err != nil{
		log.Fatal(err.Error())
	}
}
