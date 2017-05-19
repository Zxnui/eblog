package model

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"time"
)

var(
	RedisClient *redis.Pool
)

func init() {
	RedisClient = &redis.Pool{
		MaxIdle: beego.AppConfig.DefaultInt("reids_maxId",1),
		MaxActive: beego.AppConfig.DefaultInt("redis_active",10),
		IdleTimeout: time.Duration(beego.AppConfig.DefaultInt("redis_timeout",60)) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redis_url"))
			if err != nil {
				return nil, err
			}

			// auth password
			if _,err:=c.Do("AUTH",beego.AppConfig.String("redis_pwd"));err!=nil{
				c.Close()
				return nil,err
			}

			// select db
			if _, err := c.Do("SELECT", beego.AppConfig.DefaultInt("redis_db",0)); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}