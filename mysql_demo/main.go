package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_demo"
	//这里需要去初始化，而不是去新声明一个变量
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//try to connect to database(check if dsn is right)
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to database failed,err:%v\n", err)
		return
	}
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
	}
	defer db.Close()
	//ensure it is under the error jurge
	//return the resource
	fmt.Println("connect to db...")
}
