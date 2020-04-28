package assets

import (
	"github.com/garyburd/redigo/redis"
	cf "myframe/conf"
)

var Redis redis.Conn

func init() {
	var err error
	conf,err:=cf.GetArea("redis")
	if err != nil {
		return
	}
	Redis,err = redis.Dial("tcp",conf["REDIS_HOST"]+":"+conf["REDIS_PORT"])
	if err != nil {
		return
	}
	return
}

func Close() {
	if Redis != nil {
		Redis.Close()
	}
}