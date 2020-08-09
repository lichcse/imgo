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
	URL         string `yaml:"url"`
	Database    string `yaml:"database"`
	PoolLimit   int    `yaml:"pool_limit"`
	MaxLifetime int    `yaml:"max_life_time"`
}

// RedisConfig struct
type RedisConfig struct {
	URL string `yaml:"url"`
}

// RabbitMQConfig struct
type RabbitMQConfig struct {
	URL string `yaml:"url"`
}

// EnvironmentConfig struct
type EnvironmentConfig struct {
	Mongo    map[string]MongoConfig `yaml:"mongo"`
	MySQL    map[string]MySQLConfig `yaml:"mysql"`
	Redis    RedisConfig            `yaml:"redis"`
	RabbitMQ RabbitMQConfig         `yaml:"rabbitmq"`
	PORT     string                 `yaml:"port"`
}

// IMConfig interface
type IMConfig interface {
	Load(args []string) error
	GetPort() string
	Mongo() map[string]MongoConfig
	MongoItem(db string) MongoConfig
	MySQL() map[string]MySQLConfig
	MySQLItem(db string) MySQLConfig
	RabbitMQ() RabbitMQConfig
}

type imConfig struct{}

var cfg = EnvironmentConfig{}
var cfgMapping = map[string]string{
	"dev":       "dev.yaml",
	"stg":       "stg.yaml",
	"prd":       "prd.yaml",
	"default":   "dev.yaml",
	"test":      "test.yaml",
	"unit_test": "unit_test.yaml",
}

// NewIMConfig func
func NewIMConfig() IMConfig {
	return &imConfig{}
}

// Load func
func (c *imConfig) Load(args []string) error {
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
func (c *imConfig) GetPort() string {
	if len(cfg.PORT) <= 0 {
		return ":8080"
	}

	return fmt.Sprintf(":%s", cfg.PORT)
}

// Mongo func
func (c *imConfig) Mongo() map[string]MongoConfig {
	return cfg.Mongo
}

// MongoItem func
func (c *imConfig) MongoItem(db string) MongoConfig {
	if dbInfo, ok := cfg.Mongo[db]; ok {
		return dbInfo
	}

	return MongoConfig{}
}

// MySQL func
func (c *imConfig) MySQL() map[string]MySQLConfig {
	return cfg.MySQL
}

// MySQLItem func
func (c *imConfig) MySQLItem(db string) MySQLConfig {
	if dbInfo, ok := cfg.MySQL[db]; ok {
		return dbInfo
	}

	return MySQLConfig{}
}

// RabbitMQ func
func (c *imConfig) RabbitMQ() RabbitMQConfig {
	return cfg.RabbitMQ
}
