package config

const (
	Cluster = "cluster"
	Hosts = "hosts"
	Watch = "watch"
	Keys = "keys"
	ConfigFilePath = "./config.yaml"
)

type Config struct {
	Cluster map[string][]string
	Watch  map[string][]string
	path string
	type_ string
}

func NewConfig() *Config {
	return &Config{
		Cluster: nil,
		Watch:   nil,
		path:    ConfigFilePath,
		type_:   YAML,
	}
}

func (c *Config)GetPath() string{
	return c.path
}

func (c *Config)GetType() string {
	return c.type_
}

func (c *Config) GetHosts() []string{
	return c.Cluster[Hosts]
}

func (c *Config) GetWatchKeys() []string {
	return c.Watch[Keys]
}
