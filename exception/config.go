package exception

import "fmt"

type ConfigException struct {
	message string
	level int8
}

func NewConfigException(message string, level int8) ConfigException {
	return ConfigException{
		message: message,
		level:   level,
	}
}

func (c ConfigException)Error() string {
	return c.message
}

func (c ConfigException)String() string {
	return fmt.Sprintf("message: %s level: %d", c.message, c.level)
}