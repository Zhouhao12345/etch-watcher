package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"zhouhao.com/elevator/exception"
)

const (
	YAML = "yaml"
	YML = "yml"
	JSON = "json"
)


type AbsConfig interface {
	GetPath() string
	GetType() string
}

func fromFile(path string) ([]byte, error) {
	var (
		yamlFile []byte
		err error
	)
	filename, _ := filepath.Abs(path)
	if yamlFile, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	return yamlFile, nil
}

func YamlConfigBuild(c AbsConfig) error{
	var (
		yamlFile []byte
		err error
	)
	if yamlFile, err = fromFile(c.GetPath()); err != nil {
		return err
	}
	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}
	return nil
}

func JsonConfigBuild(c AbsConfig) error {
	var (
		jsonFile []byte
		err error
	)
	if jsonFile, err = fromFile(c.GetPath()); err != nil {
		return err
	}
	if err = json.Unmarshal(jsonFile, c); err != nil {
		return err
	}
	return nil
}

func LoadConfig(c AbsConfig) error{
	switch c.GetType() {
	case YAML:
		return YamlConfigBuild(c)
	case JSON:
		return JsonConfigBuild(c)
	default:
		return exception.NewConfigException("Not Found Config File Type", 0)
	}
}