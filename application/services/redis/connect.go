package redis

import (
	"github.com/go-redis/redis"
	"log"
	"zhouhao.com/elevator/application/services/config"
	"zhouhao.com/elevator/application/services/logger"
	rd "zhouhao.com/elevator/infrastructuration/redis"
)

type RedisProxy struct {
	*redis.Client
}

func optionBuilder(c rd.Config) *redis.Options {
	return &redis.Options{
		Network:            c.GetNetWork(),
		Addr:               c.GetAddr(),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           c.GetPassword(),
		DB:                 c.GetDB(),
		MaxRetries:         c.GetMaxretries(),
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           c.GetPoolSize(),
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
}

var con *RedisProxy

func init()  {
	con = &RedisProxy{
		rd.NewRedisClient(optionBuilder(config.RedisConfig)),
	}
	if result, err :=  con.Ping().Result(); err != nil {
		panic(err)
	} else {
		log.Printf("Redis Service Init Success: %s", result)
	}
}

func Exec(f func(con *RedisProxy) error) error {
	defer func() {
		if err := con.Close(); err != nil {
			logger.Logger.Panic(err)
		}
	}()
	return f(con)
}



