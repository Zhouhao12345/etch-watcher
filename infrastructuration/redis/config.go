package redis

import (
	"strconv"
	"zhouhao.com/elevator/infrastructuration/config"
)

const (
	Network = "network"
	Addr = "addr"
	Db = "db"
	Maxretries = "maxretries"
	Poolsize = "poolsize"
	Password = "password"
)

type Config struct {
	config.Config
	Redis map[string]string
}

func NewConfig() Config {
	return Config{
		Config: config.NewConfig(),
		Redis:  make(map[string]string),
	}
}

func convertStr2Int(s string) int64 {
	if db, err := strconv.ParseInt(s, 10, 16); err != nil {
		panic(err)
	} else {
		return db
	}
}

func (c Config) GetNetWork() string {
	return c.Redis[Network]
}

func (c Config) GetAddr() string {
	return c.Redis[Addr]
}

func (c Config) GetDB() int{
	return int(convertStr2Int(c.Redis[Db]))
}

func (c Config) GetMaxretries() int {
	return int(convertStr2Int(c.Redis[Maxretries]))
}

func (c Config) GetPoolSize() int{
	return int(convertStr2Int(c.Redis[Poolsize]))
}

func (c Config) GetPassword() string {
	return c.Redis[Password]
}