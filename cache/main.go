package cache

import (
	"os"
	"qiqiChat/util"
	"strconv"

	//"github.com/go-redis/redis"
	"github.com/garyburd/redigo/redis"
)

// RedisClient Redis缓存客户端单例
//var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
// func Redis() {
// 	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     os.Getenv("REDIS_ADDR"),
// 		Password: os.Getenv("REDIS_PW"),
// 		DB:       int(db),
// 	})

// 	_, err := client.Ping().Result()

// 	if err != nil {
// 		util.Err.Fatalln("连接Redis不成功", err)
// 	}

// 	RedisClient = client
// }
var RedisClient redis.Conn

func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	var options []redis.DialOption
	options = append(options, redis.DialPassword(os.Getenv("REDIS_PW")), redis.DialDatabase(int(db)))
	conn, err := redis.Dial("tcp", os.Getenv("REDIS_ADDR"), options...)
	if err != nil {
		util.Err.Fatalln("连接Redis不成功", err)
	}
	RedisClient = conn
}
