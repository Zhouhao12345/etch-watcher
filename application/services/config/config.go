package config

import (
	"log"
	"zhouhao.com/elevator/infrastructuration/config"
	"zhouhao.com/elevator/infrastructuration/etcd"
	"zhouhao.com/elevator/infrastructuration/logger"
	"zhouhao.com/elevator/infrastructuration/redis"
)

var (
	RedisConfig redis.Config = redis.NewConfig()
	EtcdConfig etcd.Config = etcd.NewConfig()
	LoggerConfig logger.Config = logger.NewConfig()
)

func initAllConfig(configs ...config.AbsConfig) error {
	for _, c := range configs {
		if err := config.LoadConfig(c); err != nil {
			return err
		}
	}
	return nil
}

func init()  {
	configs := []config.AbsConfig{
		RedisConfig,
		EtcdConfig,
		LoggerConfig,
	}
	if err := initAllConfig(configs...); err != nil {
		panic(err)
	} else {
		log.Print("Config Service Init Success")
	}
}
