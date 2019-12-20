package etcd

import "zhouhao.com/elevator/infrastructuration/config"

const (
	Hosts = "hosts"
	Watch = "watch"
	Keys = "keys"
	Key = "key"
	Privilege = "privilege"
)


type Config struct {
	config.Config
	Etcdcluster map[string][]string
	Etcdwatcher  map[string][]map[string]string
}

func NewConfig() Config {
	return Config{
		Config:      config.NewConfig(),
		Etcdcluster: make(map[string][]string),
		Etcdwatcher: make(map[string][]map[string]string),
	}
}

func (c Config) GetHosts() []string{
	return c.Etcdcluster[Hosts]
}

func (c Config) GetWatchKeys() []map[string]string {
	return c.Etcdwatcher[Keys]
}

