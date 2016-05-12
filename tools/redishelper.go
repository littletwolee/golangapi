package tools

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
)

var (
	MAX_POOL_SIZE = 20
	redisPoll chan redis.Conn
	redisnetwork = beego.AppConfig.String("redisnetwork")
	redishost = beego.AppConfig.String("redishost")
	redisport = beego.AppConfig.String("redisport")
	redispwd = beego.AppConfig.String("redispwd")
	pool = newPool()
)

type RedisHelper struct{}

func newPool() *redis.Pool {
	return &redis.Pool{
	    MaxIdle: 80,
	    MaxActive: 12000, // max number of connections
	    Dial: func() (redis.Conn, error) {
		    c, err := redis.Dial(redisnetwork,
			    redishost + ":" + redisport,
			    redis.DialPassword(redispwd))
		    if err != nil {
			    panic(err.Error())
		    }
		    return c, err
	    },
	}
}

func (r *RedisHelper) SetKVBySETEX(key interface{}, value interface{}, time int) error {
	rc := pool.Get()
	defer rc.Close()
	if _, err := redis.String(rc.Do("SETEX", key, time, value)); err != nil {
		return err
	}
	return nil
}

func (r *RedisHelper) SetKV(key interface{}, value interface{}) error {
	rc := pool.Get()
	defer rc.Close()
	if _, err := redis.String(rc.Do("SET", key, value)); err != nil {
		return err
	}
	return nil 
}

func (r *RedisHelper) GetVByK(key interface{}, returntype string) (value interface{}, err error) {
	rc := pool.Get()
	defer rc.Close()
	switch returntype {
	case "string":
		return redis.String(rc.Do("GET", key))
	case "int":
		return redis.Int(rc.Do("GET", key))
	case "bool":
		return redis.Bool(rc.Do("GET", key))
	case "bytes":
		return redis.Bytes(rc.Do("GET", key))
	}
	return value, nil
}
