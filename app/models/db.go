package models

import (
	"github.com/gocql/gocql"
	"fmt"
)


/*
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
*/

func getDB() (*gocql.Session,error) {
	fmt.Println("here is a new connection to cluster!!!!!")
	cluster := gocql.NewCluster("127.0.0.1" )
	cluster.Keyspace = "firstdb"
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
	if err!=nil {
		fmt.Println(err.Error())
	}
	return session, err
}

