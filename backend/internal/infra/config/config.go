package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func New(appName string) (*Config, error) {
	fileConfig := getConfigFile(appName)

	f, err := os.Open(fileConfig)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	cfg.AppName = appName
	return &cfg, nil
}

func getConfigFile(appName string) string {
	var (
		filename = fmt.Sprintf("%s/config.yaml", appName)
	)

	dir, _ := os.Getwd()

	return filepath.Join(dir, "files/etc", filename)
}

type Config struct {
	AppName  string
	Env      string   `yaml:"env"`
	Server   Server   `yaml:"server"`
	Database DBConfig `yaml:"database"`
}

type Server struct {
	Port         int32 `yaml:"port"`
	WriteTimeout int64 `yaml:"write_timeout"`
	ReadTimeout  int64 `yaml:"read_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
