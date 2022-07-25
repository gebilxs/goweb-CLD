package main

import (
	"net/http"
	"time"

	ratelimit2 "github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
	ratelimit1 "go.uber.org/ratelimit" //漏桶 中间件
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func heiHandler(c *gin.Context) {
	c.String(http.StatusOK, "ha")
}

//基于漏桶的中间件1
func rateLimit1() func(ctx *gin.Context) {
	//生成一个限流器
	rl := ratelimit1.New(100)
	return func(c *gin.Context) {
		//取水滴
		if rl.Take().Sub(time.Now()) > 0 {
			//time.Sleep(rl.Take().Sub(time.Now())) 需要等待这么长的时间水滴才会落下
			c.String(http.StatusOK, "ratelimit")
			c.Abort()
			return
		}
		c.Next()
	}
}

//基于令牌桶的中间件1
func rateLimit2(fillInterval time.Duration, cap int64) func(ctx *gin.Context) {
	rl := ratelimit2.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		//rl.Take()   这一次取令牌可以欠账
		//rl.TakeAvaliable()  有令牌的时候才能够取到令牌
		if rl.TakeAvailable(1) != 1 {
			c.String(http.StatusOK, "ratelimit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
func main() {
	r := gin.Default()

	r.GET("/ping", rateLimit1(), pingHandler)
	r.GET("hei", rateLimit2(2*time.Second, 1), heiHandler)

	r.Run()
}
