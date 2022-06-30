package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//声明一个全局的rdb变量
var rdb *redis.Client

//初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  //no password set
		DB:       0,   //use default DB
		PoolSize: 100, //连接池大小
	})

	_, err = rdb.Ping().Result()
	return err
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed,err:%v\n", err)
		return
	}
	fmt.Println("connect redis success...")
	//回收
	defer rdb.Close()
}
