package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), //no password set
		DB:       viper.GetInt("redis.db"),          //use default DB                                 //use default DB
		PoolSize: viper.GetInt("redis.pool_size"),   //连接池大小
	})

	_, err = rdb.Ping().Result()
	return err
}
func Close() {
	_ = rdb.Close()
}
