package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_demo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//ensure it is under the error jurge
	//return the resource

	//try to connect to database(check if dsn is right)
	err = db.Ping()
	if err != nil {
		fmt.Println("connect to database failed,err:", err)
		return
	}
	fmt.Println("connect to db...")
}
