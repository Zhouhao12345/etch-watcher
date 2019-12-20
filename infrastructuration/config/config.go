package config

const (
	Redis = "redis"
	ConfigFilePath = "./config.yaml"
)

type Config struct {
	path string
	type_ string
}

func NewConfig() Config {
	return Config{
		path:        ConfigFilePath,
		type_:       YAML,
	}
}

func (c Config)GetPath() string{
	return c.path
}

func (c Config)GetType() string {
	return c.type_
}


