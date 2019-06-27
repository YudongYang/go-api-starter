package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	RedisCache       cache.Cache
	Redis_Cache_Host string
	Redis_Cache_Db   string
	Redis_Cache_Key  string
)

// 初始化缓存
func init() {
	Redis_Cache_Host = beego.AppConfig.String("cache.host")
	Redis_Cache_Db = beego.AppConfig.String("cache.db")
	Redis_Cache_Key = beego.AppConfig.String("cache.key")
	Redis_CachePassword := beego.AppConfig.String("cache.password")
	RedisCache, _ = cache.NewCache("redis", `{"key":"`+Redis_Cache_Key+`","conn":"`+Redis_Cache_Host+`","dbNum":"`+Redis_Cache_Db+`","password":"`+Redis_CachePassword+`"}`)
	Logger.Info("Init: RedisCache: %s", Redis_Cache_Key)
}
