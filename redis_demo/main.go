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
func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
	//rdb.HGetAll()
	//rdb.HGet()
	//rdb.HMGet()
}
func hgetDemo() {
	v, err := rdb.HGetAll("user").Result()
	if err != nil {
		//redis.Nil(err)
		fmt.Printf("hgetall failed,err:%v\n", err)
		return
	}
	fmt.Println(v)

	v1 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println(v1)

	v2 := rdb.HGet("user", "age").Val()
	fmt.Println(v2)
}
func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed,err:%v\n", err)
		return
	}
	fmt.Println("connect redis success...")
	//回收
	defer rdb.Close()

	//redisExample
	//hgetDemo()
	redisExample2()
}
