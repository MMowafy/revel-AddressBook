package models

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"

)

func getDB()(*sql.DB, error)   {
	var  err error
	var db  *sql.DB
	fmt.Println("here is a new connection to db!!!!!")
	db, err = sql.Open("mysql", "root:mowafymysql@/firstdb")
	if err!=nil {
		fmt.Println(err.Error())
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
	}
	return db, err
}
