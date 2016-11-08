package hippo

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Env      string
	Database DatabaseConfig
}

type ConfigFile struct {
	Database map[string]DatabaseConfig `yaml:"database"`
}

func NewConfig(env string, path string) *Config {
	filename, _ := filepath.Abs(path)
	yamlData, err := ioutil.ReadFile(filename)

	checkError(err)

	var configFile ConfigFile
	err = yaml.Unmarshal(yamlData, &configFile)

	checkError(err)

	config := Config{Env: env}
	config.Database = configFile.Database[env]

	return &config
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
