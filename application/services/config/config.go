package config

import (
	"zhouhao.com/elevator/infrastructuration/config"
	"zhouhao.com/elevator/infrastructuration/etcd"
	"zhouhao.com/elevator/infrastructuration/redis"
)

var (
	RedisConfig redis.Config = redis.NewConfig()
	EtcdConfig etcd.Config = etcd.NewConfig()
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
	}
	if err := initAllConfig(configs...); err != nil {
		panic(err)
	}
}
