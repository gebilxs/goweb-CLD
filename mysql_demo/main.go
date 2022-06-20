package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
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

	//数值根据你的业务情况发生变化
	db.SetConnMaxIdleTime(time.Second * 10)
	db.SetMaxOpenConns(200)   //最大连接数
	db.SetConnMaxIdleTime(60) //最大空闲连接数
	return
}

//查询单条数据
type user struct {
	id   int
	age  int
	name string
}

func queryRowDemo() {
	sqlStr := "select id,name,age from user where id=?"
	var u user
	//非常重要，确保QueryRow之后调用Scan方法，否则持有的数据库不会被释放

	//row := db.QueryRow(sqlStr, 1)
	//err := row.Scan(&u.id, &u.name, &u.age)
	//优化--采用链式调用
	err := db.QueryRow(sqlStr, 3).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
	}
	defer db.Close()
	//ensure it is under the error jurge
	//return the resource
	fmt.Println("connect to db...")
	queryRowDemo()

}
