package utils

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

var (
	// 定义常量
	// 默认的链接池
	RedisPool *redis.Pool
	// 自定义的链接池
	RedisPoolSession *redis.Pool
	REDIS_HOST       string
	REDIS_DB         int
	REDIS_DB_Session int
	REDIS_PASSWORD   string
)

func initDb(db int, redisPool *redis.Pool) {
	// 从配置文件获取redis的ip
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_PASSWORD = beego.AppConfig.String("redis.password")
	// 建立连接池
	redisPool = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			option := redis.DialPassword(REDIS_PASSWORD)
			c, err := redis.Dial("tcp", REDIS_HOST, option)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", db)
			return c, nil
		},
	}
}

// 初始化 Redis 链接池
// 从 RedisPool 中 获取 Redis 链接后，记得 defer close 掉
func init() {
	REDIS_DB = beego.AppConfig.DefaultInt("redis.db", 0)
	REDIS_DB_Session = beego.AppConfig.DefaultInt("redis.db_session", 1)
	initDb(REDIS_DB, RedisPool)
	initDb(REDIS_DB_Session, RedisPoolSession)
	Logger.Info("Init: RedisPool: %s", RedisPool)
	Logger.Info("Init: RedisPoolSession: %s", RedisPoolSession)
}
