package mysql

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	//这里需要去初始化，而不是去新声明一个变量
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect to DB failed,err:%v\n", zap.Error(err))
		return
	}

	//try to connect to database(check if dsn is right)
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to database failed,err:%v\n", err)
		return
	}
	//数值根据你的业务情况发生变化
	db.SetConnMaxIdleTime(time.Second * 10)
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) //最大连接数
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) //最大空闲连接数
	return
}
func Close() {
	_ = db.Close()
}
