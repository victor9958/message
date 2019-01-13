package model

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
)

var MyRedis cache.Cache
func init(){
	cacheconn := beego.AppConfig.String("redis.conn")
	cachepwd := beego.AppConfig.String("redis.pwd")
	cachedb := beego.AppConfig.String("redis.cachedbname")
	cachestr := `{"key":"victor","conn":"`+cacheconn+`","dbNum":"`+cachedb+`","password":"`+ cachepwd+`"}`
	beego.Info(cachestr)
	var err error
	MyRedis,err =cache.NewCache("redis",cachestr) // 制造一个redis 连接
	if err != nil{
		beego.Info("redis作为缓存错误是:",err)
	}
}

