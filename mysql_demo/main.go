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
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

//查询多条记录
func queryMultiRowDemo() {
	//判断此条记录不为0的情况
	sqlStr := "select id,name,age from user where id >?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	//非常重要，关闭rows,释放数据库连接
	defer rows.Close()

	//循环读取结果
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func insertRowDemo() {
	sqlStr := "insert into user(name,age) values (?,?)"
	ret, err := db.Exec(sqlStr, "qcy", 22)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	var theID int64
	theID, err = ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("getlastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", theID)
}

//更新数据
func updataRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 30, 1)
	if err != nil {
		fmt.Printf("updata failed,err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("updata success,affected rows:%d\n", n)
}

//删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id =?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected() //影响操作的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows:%d\n", n)

}

//循环删
//
//func deleteall() {
//	sqlStr := "delete from user where id >?"
//	ret, err:=db.Exec(sqlStr,4)
//	if err!=nil{
//		fmt.Printf("delete all failed,err:%v\n",err)
//		return
//	}
//
//}
func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
	}
	defer db.Close()
	//ensure it is under the error jurge
	//return the resource
	fmt.Println("connect to db...")
	deleteRowDemo()
	updataRowDemo()
	queryRowDemo()
	insertRowDemo()
	queryMultiRowDemo()

}
