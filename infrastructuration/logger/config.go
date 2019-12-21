package logger

import "zhouhao.com/elevator/infrastructuration/config"

const (
	Filepath = "filepath"
)


type Config struct {
	config.Config
	Logger map[string]string
}

func NewConfig() Config {
	return Config{
		Config:      config.NewConfig(),
		Logger: make(map[string]string),
	}
}

func (c Config) GetFilePath() string{
	return c.Logger[Filepath]
}


