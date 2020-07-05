package utils

import (
	"errors"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// MongoConfig struct
type MongoConfig struct {
	URL       string `yaml:"url"`
	Database  string `yaml:"database"`
	PoolLimit int    `yaml:"pool_limit"`
}

// MySQLConfig struct
type MySQLConfig struct {
	URL       string `yaml:"url"`
	Database  string `yaml:"database"`
	PoolLimit int    `yaml:"pool_limit"`
}

// RedisConfig struct
type RedisConfig struct {
	URL string `yaml:"url"`
}

// EnvironmentConfig struct
type EnvironmentConfig struct {
	Mongo map[string]MongoConfig `yaml:"mongo"`
	MySQL map[string]MySQLConfig `yaml:"mysql"`
	Redis RedisConfig            `yaml:"redis"`
	PORT  string                 `yaml:"port"`
}

// IMConfig interface
type IMConfig interface {
	Load(args []string) error
	GetPort() string
	Mongo() map[string]MongoConfig
	MongoItem(db string) MongoConfig
	MySQL() map[string]MySQLConfig
	MySQLItem(db string) MySQLConfig
}

type imConfig struct{}

var cfg = EnvironmentConfig{}
var cfgMapping = map[string]string{
	"test":      "test.yaml",
	"dev":       "development.yaml",
	"stg":       "staging.yaml",
	"prod":      "production.yaml",
	"default":   "development.yaml",
	"unit_test": "unit_test.yaml",
}

// NewIMConfig func
func NewIMConfig() IMConfig {
	return &imConfig{}
}

// Load func
func (m *imConfig) Load(args []string) error {
	if len(args) < 1 {
		return errors.New("")
	}

	env, ok := cfgMapping[args[0]]
	if !ok {
		env = cfgMapping["default"]
	}

	configData, err := ioutil.ReadFile("./config/" + env)
	if err != nil {
		return err
	}

	cfg = EnvironmentConfig{}
	return yaml.Unmarshal(configData, &cfg)
}

// GetPort func
func (m *imConfig) GetPort() string {
	if len(cfg.PORT) <= 0 {
		return ":8080"
	}

	return fmt.Sprintf(":%s", cfg.PORT)
}

// Mongo func
func (m *imConfig) Mongo() map[string]MongoConfig {
	return cfg.Mongo
}

// MongoItem func
func (m *imConfig) MongoItem(db string) MongoConfig {
	if dbInfo, ok := cfg.Mongo[db]; ok {
		return dbInfo
	}

	return MongoConfig{}
}

// MySQL func
func (m *imConfig) MySQL() map[string]MySQLConfig {
	return cfg.MySQL
}

// MySQLItem func
func (m *imConfig) MySQLItem(db string) MySQLConfig {
	if dbInfo, ok := cfg.MySQL[db]; ok {
		return dbInfo
	}

	return MySQLConfig{}
}
