package redis

import (
	"fmt"
	"goweb-CLD/web_CLD/setting"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			//viper.GetString("redis.host"),
			//viper.GetInt("redis.port"),
			cfg.Host,
			cfg.Port,
		),
		//Password: viper.GetString("redis.password"), //no password set
		//DB:       viper.GetInt("redis.db"),          //use default DB                                 //use default DB
		//PoolSize: viper.GetInt("redis.pool_size"),   //连接池大小
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return err
}
func Close() {
	_ = rdb.Close()
}
